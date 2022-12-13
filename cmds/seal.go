package cmds

type SealCommand struct {
	CreateAccount         CreateAccountCommand         `cmd:"" name:"create-account" help:"create new account"`
	KeyUpdater            KeyUpdaterCommand            `cmd:"" name:"key-updater" help:"update account keys"`
	Transfer              TransferCommand              `cmd:"" name:"transfer" help:"transfer"`
	CurrencyRegister      CurrencyRegisterCommand      `cmd:"" name:"currency-register" help:"register new currency"`
	CurrencyPolicyUpdater CurrencyPolicyUpdaterCommand `cmd:"" name:"currency-policy-updater" help:"update currency policy"`
}

func NewSealCommand() SealCommand {
	return SealCommand{
		CreateAccount:         NewCreateAccountCommand(),
		KeyUpdater:            NewKeyUpdaterCommand(),
		Transfer:              NewTransferCommand(),
		CurrencyRegister:      NewCurrencyRegisterCommand(),
		CurrencyPolicyUpdater: NewCurrencyPolicyUpdaterCommand(),
	}
}
