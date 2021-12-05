package imports

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
	"golang.org/x/tools/internal/gocommand"
)

type importsLocalTestCase struct {
	options *Options

	imports        []string
	inputTpl       string
	ref            string
	shuffleImports bool
}

func TestImportsLocal(t *testing.T) {
	imports := []string{
		`"bytes"`,
		`"database/sql"`,
		`"fmt"`,
		`"strconv"`,
		`"time"`,
		`"github.com/lib/pq"`,
		`"github.com/pkg/errors"`,
		`"gopkg.in/gorp.v1"`,
		`"other-repo.by/X/Y"`,
		`nu "other-repo.by/X/NotUsed"`,
		`"gitlab.internal.com/X/Y/datastructures"`,
		`a "gitlab.internal.com/X/Y/A"`,
		`b "gitlab.internal.com/X/Y/B"`,
		`c "gitlab.internal.com/X/Y/C"`,
		`a1 "project-local-package-1/X/Y/A"`,
		`b1 "project-local-package-1/X/Y/B"`,
		`"project-local-package-2/X/Y/A"`,
		`"project-local-package-2/X/Y/B"`,
		`"project-local-package-2/X/Y/NotUsed"`,
	}

	inputTpl := `package pkg

import (
%s
)

var _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bytes.Buffer{}, sql.Stmt{}, fmt.Sprint(), strconv.IntSize, time.February, pq.Error{}, errors.StackTrace{}, gorp.DbMap{}, Y.F, datastructures.F, a.F, b.F, c.F, a1.F, b1.F, A.F, B.F
`

	ref := `package pkg

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/gorp.v1"
	"other-repo.by/X/Y"

	a "gitlab.internal.com/X/Y/A"
	b "gitlab.internal.com/X/Y/B"
	c "gitlab.internal.com/X/Y/C"
	"gitlab.internal.com/X/Y/datastructures"

	"project-local-package-2/X/Y/A"
	"project-local-package-2/X/Y/B"

	a1 "project-local-package-1/X/Y/A"
	b1 "project-local-package-1/X/Y/B"
)

var _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bytes.Buffer{}, sql.Stmt{}, fmt.Sprint(), strconv.IntSize, time.February, pq.Error{}, errors.StackTrace{}, gorp.DbMap{}, Y.F, datastructures.F, a.F, b.F, c.F, a1.F, b1.F, A.F, B.F
`

	options := Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
		Env: &ProcessEnv{
			GocmdRunner: &gocommand.Runner{},
		},
		LocalPrefix: "project-local-package-notexist,gitlab.internal.com,project-local-package-notexist2,project-local-package-2,project-local-package-1",
	}

	//-------------------

	N := 1000

	testcases := make([]importsLocalTestCase, N)

	testcases[0] = importsLocalTestCase{
		options: &options,

		imports:        imports,
		inputTpl:       inputTpl,
		ref:            ref,
		shuffleImports: false,
	}

	for i := 1; i < N; i++ {
		testcases[i] = importsLocalTestCase{
			options: &options,

			imports:        imports,
			inputTpl:       inputTpl,
			ref:            ref,
			shuffleImports: true,
		}
	}

	for _, tc := range testcases {
		err := tc.assertProcessEquals()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func (tc *importsLocalTestCase) prepareInput() string {
	var imps bytes.Buffer
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if tc.shuffleImports {
		r.Shuffle(len(tc.imports), func(i int, j int) {
			tc.imports[i], tc.imports[j] = tc.imports[j], tc.imports[i]
		})
	}

	//insert random empty lines
	for i := 0; i < len(tc.imports); i++ {
		if r.Int()%2 == 0 {
			imps.WriteString("\n")
			if r.Int()%2 == 0 {
				imps.WriteString("\n")
			}
		}
		imps.WriteString(tc.imports[i])
		imps.WriteString("\n")
	}

	return fmt.Sprintf(tc.inputTpl, imps.String())
}

func (tc *importsLocalTestCase) assertProcessEquals() error {
	input := tc.prepareInput()
	b, err := Process("", []byte(input), tc.options)
	if err != nil {
		return err
	}

	if string(b) != tc.ref {
		return fmt.Errorf("Not equal!\n%s", string(b))
	}

	return nil
}
