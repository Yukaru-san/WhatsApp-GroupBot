package wabot

import (
	"fmt"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

type cmd struct{}

var (
	consoleWriteTo = ""

	contacList []whatsapp.Contact

	session whatsapp.Session
	conn    *whatsapp.Conn

	longConnName  = "Whatsapp-GroupBot"
	shortConnName = "wabot"

	startTime        = uint64(time.Now().Unix())
	errorTimeout     = time.Minute * 1
	enableAutosaving = true
	autosaveInterval = time.Minute * 3

	encrypKey        = []byte("r4gyXrWSPXzvpBZJ")
	showTextMessages = true
	useContactName   = false

	useNicknames       = false
	nicknameUpdateText = "Your nickname has been updated!"

	qrCodeFile      = "qr.png"
	sessionFile     = "storedSession.dat"
	usersFile       = "storedUsers.dat"
	whatsappTimeout = 20 * time.Second
)

// SetEncryptionKey replaces the standard encryption Key
//  - Key has to be 16 Byte long
//  - Returns false if the key is bad
func SetEncryptionKey(key []byte) bool {
	if len(key) < 16 || len(key) > 16 {
		return false
	}

	encrypKey = key
	return true
}

// SetSessionFilePath changes the name a file should be saved in ([folder/]filename)
func SetSessionFilePath(path string) {
	sessionFile = path
}

// SetUsersFilePath changes the location the users will be saved in ([folder/]filename)
func SetUsersFilePath(path string) {
	usersFile = path
}

// SetQRFilePath changes the location the users will be saved in ([folder/]filename)
func SetQRFilePath(path string) {
	qrCodeFile = path
}

// UseContactNames tells the bot to use names saved in contacts (or not)
func UseContactNames(use bool) {
	useContactName = use
}

// DeactivateAutoSaving disables automatic saving
func DeactivateAutoSaving() {
	enableAutosaving = false
}

// SetAutosaveInterval - interval of userdata saving
func SetAutosaveInterval(interval time.Duration) {
	autosaveInterval = interval
}

// SetErrorTimeout sets the default time to reconnect after
// an error caused the program to disconnect
func SetErrorTimeout(timeout time.Duration) {
	errorTimeout = timeout
}

// AddTextCommand adds a command that only works in all groups
func AddTextCommand(cmd string, functionToExecute func(whatsapp.TextMessage)) {
	commands = append(commands, Command{prefix: cmd, function: functionToExecute})
}

// AddGroupCommand adds a command that only works on certain groups
func AddGroupCommand(cmd string, groupsToWorkIn []string, functionToExecute func(whatsapp.TextMessage)) {

	groupIDList := []string{}

	for _, g := range groupsToWorkIn {
		fmt.Println("g:", g, "NameToJid:", JidToGroupID(NameToJid(g)))
		groupIDList = append(groupIDList, JidToGroupID(NameToJid(g)))
	}

	commands = append(commands, Command{prefix: cmd, groups: groupIDList, function: functionToExecute})

	//	commands = append(commands, Command{prefix: cmd, groups: groupsToWorkIn, function: functionToExecute})
}

// SetImageHandler calls the given function when receiving an img
func SetImageHandler(functionToExecute func(whatsapp.ImageMessage)) {
	imageHandleFunction = functionToExecute
}

// SetStickerHandler calls the given function when receiving an img
func SetStickerHandler(functionToExecute func(whatsapp.StickerMessage)) {
	stickerHandleFunction = functionToExecute
}

// SetDefaultTextHandleFunction calls the given function when no commands have fit
func SetDefaultTextHandleFunction(functionToExecute func(whatsapp.TextMessage)) {
	defaultTextHandleFunction = functionToExecute
}

// DisplayTextMessagesInConsole toggles visibility in console
func DisplayTextMessagesInConsole(display bool) {
	showTextMessages = display
}

// HandleError prints potential errors (unused)
func (cmd) HandleError(err error) {
}

// HandleContactList fills the contacList on load
func (cmd) HandleContactList(contacts []whatsapp.Contact) {
	for _, c := range contacts {
		contacList = append(contacts, c)
	}
}

// SetNicknameUseage changes the users abbility to use nicknames (false by default) - string will be the output. "" for no output
func SetNicknameUseage(allowNicknames bool, output string) {
	useNicknames = allowNicknames
	nicknameUpdateText = output
}
