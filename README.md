# Fork of `Go Tools` to tune `goimports` behavior

This fork is created to fix next issues

- work with 3rd-party packages defined in `local` flag in more strict way
	- if few packages are provided in `local`, the each group will be intended with new line
	- groups ordering is the same as provided in `local`
- ignore user defined empty lines between imports. It allows to receive more predictable ordering

## Download/Install

```bash
$ git clone github.com/nawa/go-tools
$ cd go-tools/cmd/goimports
$ go install
```

## Comparison with standard `goimports`

The next command is used

```bash
$ goimports -local="gitlab.internal.com,project-local-package-2,project-local-package-1"
```

### Comparison of 3rd-party packages indentation and ordering:
Input:

```go
package pkg

import (
	"project-local-package-2/X/Y/B"
	"other-repo.by/X/Y"
	"project-local-package-2/X/Y/NotUsed"
	"fmt"
	"gitlab.internal.com/X/Y/datastructures"
	"database/sql"
	c "gitlab.internal.com/X/Y/C"
	"gopkg.in/gorp.v1"
	"bytes"
	a "gitlab.internal.com/X/Y/A"
	b "gitlab.internal.com/X/Y/B"
	b1 "project-local-package-1/X/Y/B"
	"github.com/lib/pq"
	a1 "project-local-package-1/X/Y/A"
	"strconv"
	"time"
	nu "other-repo.by/X/NotUsed"
	"project-local-package-2/X/Y/A"
	"github.com/pkg/errors"
)

var _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bytes.Buffer{}, sql.Stmt{}, fmt.Sprint(), strconv.IntSize, time.February, pq.Error{}, errors.StackTrace{}, gorp.DbMap{}, Y.F, datastructures.F, a.F, b.F, c.F, a1.F, b1.F, A.F, B.F
```

Standard `goimports`:

```go
package pkg

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
	a1 "project-local-package-1/X/Y/A"
	b1 "project-local-package-1/X/Y/B"
	"project-local-package-2/X/Y/A"
	"project-local-package-2/X/Y/B"
)

var _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bytes.Buffer{}, sql.Stmt{}, fmt.Sprint(), strconv.IntSize, time.February, pq.Error{}, errors.StackTrace{}, gorp.DbMap{}, Y.F, datastructures.F, a.F, b.F, c.F, a1.F, b1.F, A.F, B.F
```

Tuned `goimports`:

```go
package pkg

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
```

### Comparison of ignoring user defined empty lines

Input:

```go
package pkg

import (
	"project-local-package-2/X/Y/B"
	"other-repo.by/X/Y"
	"project-local-package-2/X/Y/NotUsed"
	"fmt"
	"gitlab.internal.com/X/Y/datastructures"

	"database/sql"

	c "gitlab.internal.com/X/Y/C"

	"gopkg.in/gorp.v1"

	"bytes"

	a "gitlab.internal.com/X/Y/A"

	b "gitlab.internal.com/X/Y/B"

	b1 "project-local-package-1/X/Y/B"

	"github.com/lib/pq"
	a1 "project-local-package-1/X/Y/A"

	"strconv"

	"time"
	nu "other-repo.by/X/NotUsed"

	"project-local-package-2/X/Y/A"

	"github.com/pkg/errors"
)

var _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bytes.Buffer{}, sql.Stmt{}, fmt.Sprint(), strconv.IntSize, time.February, pq.Error{}, errors.StackTrace{}, gorp.DbMap{}, Y.F, datastructures.F, a.F, b.F, c.F, a1.F, b1.F, A.F, B.F
```

Standard `goimports`:

```go
package pkg

import (
	"fmt"

	"other-repo.by/X/Y"

	"gitlab.internal.com/X/Y/datastructures"
	"project-local-package-2/X/Y/B"

	"database/sql"

	c "gitlab.internal.com/X/Y/C"

	"gopkg.in/gorp.v1"

	"bytes"

	a "gitlab.internal.com/X/Y/A"

	b "gitlab.internal.com/X/Y/B"

	b1 "project-local-package-1/X/Y/B"

	"github.com/lib/pq"

	a1 "project-local-package-1/X/Y/A"

	"strconv"

	"time"

	"project-local-package-2/X/Y/A"

	"github.com/pkg/errors"
)

var _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bytes.Buffer{}, sql.Stmt{}, fmt.Sprint(), strconv.IntSize, time.February, pq.Error{}, errors.StackTrace{}, gorp.DbMap{}, Y.F, datastructures.F, a.F, b.F, c.F, a1.F, b1.F, A.F, B.F
```

Tuned `goimports`:

```go
package pkg

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
```

# Original README

# Go Tools

[![PkgGoDev](https://pkg.go.dev/badge/golang.org/x/tools)](https://pkg.go.dev/golang.org/x/tools)

This subrepository holds the source for various packages and tools that support
the Go programming language.

Some of the tools, `godoc` and `vet` for example, are included in binary Go
distributions.

Others, including the Go `guru` and the test coverage tool, can be fetched with
`go install`.

Packages include a type-checker for Go and an implementation of the
Static Single Assignment form (SSA) representation for Go programs.

## Download/Install

The easiest way to install is to run `go install golang.org/x/tools/...@latest`.

## JS/CSS Formatting

This repository uses [prettier](https://prettier.io/) to format JS and CSS files.

The version of `prettier` used is 1.18.2.

It is encouraged that all JS and CSS code be run through this before submitting
a change. However, it is not a strict requirement enforced by CI.

## Report Issues / Send Patches

This repository uses Gerrit for code changes. To learn how to submit changes to
this repository, see https://golang.org/doc/contribute.html.

The main issue tracker for the tools repository is located at
https://github.com/golang/go/issues. Prefix your issue with "x/tools/(your
subdir):" in the subject line, so it is easy to find.
