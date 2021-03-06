// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package multipart implements MIME multipart parsing, as defined in RFC 2046.
//
// The implementation is sufficient for HTTP (RFC 2388) and the multipart bodies
// generated by popular browsers.

// multipart实现了MIME的multipart解析，参见RFC 2046。该实现适用于HTTP（RFC
// 2388）和常见浏览器生成的multipart主体。
package multipart

// File is an interface to access the file part of a multipart message. Its
// contents may be either stored in memory or on disk. If stored on disk, the
// File's underlying concrete type will be an *os.File.

// File是一个接口，实现了对一个multipart信息中文件记录的访问。它的内容可以保持在内存或者硬盘中，如果保持在硬盘中，底层类型就会是*os.File。
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

// A FileHeader describes a file part of a multipart request.

// FileHeader描述一个multipart请求的（一个）文件记录的信息。
type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	// contains filtered or unexported fields
}

// Open opens and returns the FileHeader's associated File.

// Open方法打开并返回其关联的文件。
func (fh *FileHeader) Open() (File, error)

// Form is a parsed multipart form. Its File parts are stored either in memory or
// on disk, and are accessible via the *FileHeader's Open method. Its Value parts
// are stored as strings. Both are keyed by field name.

// Form是一个解析过的multipart表格。它的File参数部分保存在内存或者硬盘上，可以使用*FileHeader类型属性值的Open方法访问。它的Value
// 参数部分保存为字符串，两者都以属性名为键。
type Form struct {
	Value map[string][]string
	File  map[string][]*FileHeader
}

// RemoveAll removes any temporary files associated with a Form.

// 删除Form关联的所有临时文件。
func (f *Form) RemoveAll() error

// A Part represents a single part in a multipart body.

// Part代表multipart主体的单独一个记录。
type Part struct {
	// The headers of the body, if any, with the keys canonicalized
	// in the same fashion that the Go http.Request headers are.
	// For example, "foo-bar" changes case to "Foo-Bar"
	//
	// As a special case, if the "Content-Transfer-Encoding" header
	// has a value of "quoted-printable", that header is instead
	// hidden from this map and the body is transparently decoded
	// during Read calls.
	Header textproto.MIMEHeader
	// contains filtered or unexported fields
}

func (p *Part) Close() error

// FileName returns the filename parameter of the Part's Content-Disposition
// header.

// 返回Part 的Content-Disposition 头的文件名参数。
func (p *Part) FileName() string

// FormName returns the name parameter if p has a Content-Disposition of type
// "form-data". Otherwise it returns the empty string.

// 如果p的Content-Disposition头值为"form-data"，则返回名字参数；否则返回空字符串。
func (p *Part) FormName() string

// Read reads the body of a part, after its headers and before the next part (if
// any) begins.

// Read方法读取一个记录的主体，也就是其头域之后到下一记录之前的部分。
func (p *Part) Read(d []byte) (n int, err error)

// Reader is an iterator over parts in a MIME multipart body. Reader's underlying
// parser consumes its input as needed. Seeking isn't supported.

// Reader是MIME的multipart主体所有记录的迭代器。Reader的底层会根据需要解析输入，不支持Seek。
type Reader struct {
	// contains filtered or unexported fields
}

// NewReader creates a new multipart Reader reading from r using the given MIME
// boundary.
//
// The boundary is usually obtained from the "boundary" parameter of the message's
// "Content-Type" header. Use mime.ParseMediaType to parse such headers.

// 函数使用给出的MIME边界和r创建一个multipart读取器。
//
// 边界一般从信息的"Content-Type"
// 头的"boundary"属性获取。可使用mime.ParseMediaType函数解析这种头域。
func NewReader(r io.Reader, boundary string) *Reader

// NextPart returns the next part in the multipart or an error. When there are no
// more parts, the error io.EOF is returned.
func (r *Reader) NextPart() (*Part, error)

// ReadForm parses an entire multipart message whose parts have a
// Content-Disposition of "form-data". It stores up to maxMemory bytes of the file
// parts in memory and the remainder on disk in temporary files.
func (r *Reader) ReadForm(maxMemory int64) (f *Form, err error)

// A Writer generates multipart messages.

// Writer类型用于生成multipart信息。
type Writer struct {
	// contains filtered or unexported fields
}

// NewWriter returns a new multipart Writer with a random boundary, writing to w.

// NewWriter函数返回一个设定了一个随机边界的Writer，数据写入w。
func NewWriter(w io.Writer) *Writer

// Boundary returns the Writer's boundary.

// 方法返回该Writer的边界。
func (w *Writer) Boundary() string

// Close finishes the multipart message and writes the trailing boundary end line
// to the output.

// Close方法结束multipart信息，并将结尾的边界写入底层io.Writer接口。
func (w *Writer) Close() error

// CreateFormField calls CreatePart with a header using the given field name.

// CreateFormField方法使用给出的属性名调用CreatePart方法。
func (w *Writer) CreateFormField(fieldname string) (io.Writer, error)

// CreateFormFile is a convenience wrapper around CreatePart. It creates a new
// form-data header with the provided field name and file name.

// CreateFormFile是CreatePart方法的包装，
// 使用给出的属性名和文件名创建一个新的form-data头。
func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error)

// CreatePart creates a new multipart section with the provided header. The body of
// the part should be written to the returned Writer. After calling CreatePart, any
// previous part may no longer be written to.

// CreatePart方法使用提供的header创建一个新的multipart记录。该记录的主体应该写入返回的Writer接口。调用本方法后，任何之前的记录都不能再写入。
func (w *Writer) CreatePart(header textproto.MIMEHeader) (io.Writer, error)

// FormDataContentType returns the Content-Type for an HTTP multipart/form-data
// with this Writer's Boundary.

// 方法返回w对应的HTTP
// multipart请求的Content-Type的值，多以multipart/form-data起始。
func (w *Writer) FormDataContentType() string

// SetBoundary overrides the Writer's default randomly-generated boundary separator
// with an explicit value.
//
// SetBoundary must be called before any parts are created, may only contain
// certain ASCII characters, and must be 1-69 bytes long.

// SetBoundary方法重写Writer默认的随机生成的边界为提供的boundary参数。方法必须在创建任何记录之前调用，boundary只能包含特定的ascii字符，并且长度应在1-69字节之间。
func (w *Writer) SetBoundary(boundary string) error

// WriteField calls CreateFormField and then writes the given value.

// WriteField方法调用CreateFormField并写入给出的value。
func (w *Writer) WriteField(fieldname, value string) error
