
Currently there is no way for the user to query the <command>Opts struct,
nor the generated commandstring. Neither before nor after the command executed.
Here I keep ideas to make that possible. But it will not be easy to not break
the nice and simple user experience the current API offers.

Maybe I can experiment with these in a seperate cloned repo.

// ------------------------------------------------------

// Ex.: hg.Exec("identify", "/home/me/DEV/go/myrepo/", Verbose(True))
func (hgcl *HgClient) Exec(cmd string, args string, opts ...optionAdder) ([]byte, error) {
// Ex.: hg.Run("identify", "/home/me/DEV/go/myrepo/", Verbose(True))
func (hgcl *HgClient) Run(cmd string, args string, opts ...optionAdder) ([]byte, error) {

// ------------------------------------------------------

type HgCmd struct {
	cmdStr []string
	// cmdOpts: to be substituted with an identifyOpts instance f.i.
	cmdOpts /*struct*/ interface{}
}

type commandExecuter interface {
	Exec ...
}
var Identify HgCmd
func (hc *HgCmd) Exec ([]byte, error) () {}

func newHgCmd(cmdname string, hgcl *HgClient) ([]byte, error)

func newIdentifyCmd(hgcl *HgClient)


// ------------------------------------------------------

// CmdString = "getCmdString"
// (a getter in Go does not have the 'get' part in it's name)
// Use function to get the commmand details, to avoid exporting the data,
// to avoid tampering with it but via the regular ways !
func CmdString("identify") (string, error)

// Maybe a:
func [(...)] Identify(...) ([]byte, error) {}
// Is still possible, as a convenience, but then I'll have to pass it
// an instance of HgCmd I guess.

