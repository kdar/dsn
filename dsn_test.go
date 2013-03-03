package dsn

import (
  "fmt"
  "testing"
)

var dsnTests = []struct {
  in  string
  out *DSN
}{
  {"type://username:password@protocol(address)/database?param=value", &DSN{Type: "type", Username: "username", Password: "password", Protocol: "protocol", Address: "address", Database: "database", Params: map[string]string{"param": "value"}}},
  {"/", &DSN{Type: "", Username: "", Password: "", Protocol: "", Address: "", Database: "", Params: map[string]string{}}},
  {"tcp(127.0.0.1:1234)/db", &DSN{Type: "", Username: "", Password: "", Protocol: "tcp", Address: "127.0.0.1:1234", Database: "db", Params: map[string]string{}}},
  {"tcp(127.0.0.1:1234)/db?a=b,c", &DSN{Type: "", Username: "", Password: "", Protocol: "tcp", Address: "127.0.0.1:1234", Database: "db", Params: map[string]string{"a": "b,c"}}},
  {"mysql://unix(/path/to/the/socket)/db?charset=utf8", &DSN{Type: "mysql", Username: "", Password: "", Protocol: "unix", Address: "/path/to/the/socket", Database: "db", Params: map[string]string{"charset": "utf8"}}},
}

func TestParseDSN(t *testing.T) {
  var res, expect string

  for i, tt := range dsnTests {
    dsn := ParseDSN(tt.in)
    res = fmt.Sprintf("%+v", dsn)
    expect = fmt.Sprintf("%+v", tt.out)
    if res != expect {
      t.Errorf("%d. parseDSN(%q) => %q, want %q", i, tt.in, res, expect)
    }
  }
}

var dsnStringTests = []struct {
  in  *DSN
  out string
}{
  {dsnTests[0].out, dsnTests[0].in},
  {dsnTests[1].out, dsnTests[1].in},
  {dsnTests[2].out, dsnTests[2].in},
  {dsnTests[3].out, dsnTests[3].in},
  {dsnTests[4].out, dsnTests[4].in},
}

func TestDSNToString(t *testing.T) {
  var res, expect string

  for i, tt := range dsnStringTests {
    res = tt.in.String()
    expect = tt.out
    if res != expect {
      t.Errorf("%d. dsn.String(%q) => %q, want %q", i, tt.in, res, expect)
    }
  }
}
