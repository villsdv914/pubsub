package publog
import (
	"github.com/sirupsen/logrus"
	"os"
)
var Logrs = logrus.New()
func init(){
	Logrs.Out = os.Stdout
	Logrs.Formatter = &logrus.JSONFormatter{}

}
