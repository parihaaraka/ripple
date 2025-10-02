package data

// Horrible look up tables
// Could all this be one big map?

type LedgerEntryType uint16
type TransactionType uint16

const (
	// LedgerEntryType values come from rippled's "LedgerFormats.h"
	// https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types#ledger-entry-types
	ACCOUNT_ROOT                         LedgerEntryType = 0x61 // 'a'
	AMENDMENTS                           LedgerEntryType = 0x66 // 'f'
	AMM_LT                               LedgerEntryType = 0x79
	BRIDGE                               LedgerEntryType = 0x69 // 'i'
	CHECK                                LedgerEntryType = 0x43 // 'C'
	DEPOSIT_PRE_AUTH                     LedgerEntryType = 0x70 // 'p'
	DID                                  LedgerEntryType = 0x49 // 'I'
	DIRECTORY                            LedgerEntryType = 0x64 // 'd'
	ESCROW                               LedgerEntryType = 0x75 // 'u'
	FEE_SETTINGS                         LedgerEntryType = 0x73 // 's'
	LEDGER_HASHES                        LedgerEntryType = 0x68 // 'h'
	NEGATIVE_UNL                         LedgerEntryType = 0x4e // 'N'
	NFTOKEN_OFFER                        LedgerEntryType = 0x37 // '7'
	NFTOKEN_PAGE                         LedgerEntryType = 0x50 // 'P'
	OFFER                                LedgerEntryType = 0x6f // 'o'
	ORACLE                               LedgerEntryType = 0x80
	PAY_CHANNEL                          LedgerEntryType = 0x78 // 'x'
	RIPPLE_STATE                         LedgerEntryType = 0x72 // 'r'
	SIGNER_LIST                          LedgerEntryType = 0x53 // 'S'
	TICKET                               LedgerEntryType = 0x54 // 'T'
	XCHAIN_OWNED_CLAIM_ID                LedgerEntryType = 0x71 // 'q'
	XCHAIN_OWNED_CREATE_ACCOUNT_CLAIM_ID LedgerEntryType = 0x74
	CREDENTIAL                           LedgerEntryType = 0x89

	// TransactionType values come from rippled's "TxFormats.h"
	// https://xrpl.org/docs/references/protocol/transactions/types#transaction-types
	PAYMENT                               TransactionType = 0
	ESCROW_CREATE                         TransactionType = 1
	ESCROW_FINISH                         TransactionType = 2
	ACCOUNT_SET                           TransactionType = 3
	ESCROW_CANCEL                         TransactionType = 4
	SET_REGULAR_KEY                       TransactionType = 5
	OFFER_CREATE                          TransactionType = 7
	OFFER_CANCEL                          TransactionType = 8
	TICKET_CREATE                         TransactionType = 10
	SIGNER_LIST_SET                       TransactionType = 12
	PAYCHAN_CREATE                        TransactionType = 13
	PAYCHAN_FUND                          TransactionType = 14
	PAYCHAN_CLAIM                         TransactionType = 15
	CHECK_CREATE                          TransactionType = 16
	CHECK_CASH                            TransactionType = 17
	CHECK_CANCEL                          TransactionType = 18
	SET_DEPOSIT_PREAUTH                   TransactionType = 19
	TRUST_SET                             TransactionType = 20
	ACCOUNT_DELETE                        TransactionType = 21
	HOOK_SET                              TransactionType = 22
	NFTOKEN_MINT                          TransactionType = 25
	NFTOKEN_BURN                          TransactionType = 26
	NFTOKEN_CREATE_OFFER                  TransactionType = 27
	NFTOKEN_CANCEL_OFFER                  TransactionType = 28
	NFTOKEN_ACCEPT_OFFER                  TransactionType = 29
	CLAWBACK                              TransactionType = 30
	NFTOKEN_MODIFY                        TransactionType = 31
	AMM_CREATE                            TransactionType = 35
	AMM_DEPOSIT                           TransactionType = 36
	AMM_WITHDRAW                          TransactionType = 37
	AMM_VOTE                              TransactionType = 38
	AMM_BID                               TransactionType = 39
	AMM_DELETE                            TransactionType = 40
	DID_SET                               TransactionType = 41
	DID_DELETE                            TransactionType = 42
	ORACLE_SET                            TransactionType = 43
	ORACLE_DELETE                         TransactionType = 44
	XCHAIN_ACCOUNT_CREATE_COMMIT          TransactionType = 45
	XCHAIN_ADD_ACCOUNT_CREATE_ATTESTATION TransactionType = 46
	XCHAIN_ADD_CLAIM_ATTESTATION          TransactionType = 47
	XCHAIN_CLAIM                          TransactionType = 48
	XCHAIN_COMMIT                         TransactionType = 49
	XCHAIN_CREATE_BRIDGE                  TransactionType = 50
	XCHAIN_CREATE_CLAIM_ID                TransactionType = 51
	XCHAIN_MODIFY_BRIDGE                  TransactionType = 52
	CREDENTIAL_CREATE                     TransactionType = 53
	AMENDMENT                             TransactionType = 100
	SET_FEE                               TransactionType = 101
	UNL_MODIFY                            TransactionType = 102
)

var LedgerFactory = [...]func() Hashable{
	func() Hashable { return &Ledger{} },
}

var LedgerEntryFactory = [...]func() LedgerEntry{
	ACCOUNT_ROOT:          func() LedgerEntry { return &AccountRoot{leBase: leBase{LedgerEntryType: ACCOUNT_ROOT}} },
	AMENDMENTS:            func() LedgerEntry { return &Amendments{leBase: leBase{LedgerEntryType: AMENDMENTS}} },
	AMM_LT:                func() LedgerEntry { return &AMM{leBase: leBase{LedgerEntryType: AMM_LT}} },
	BRIDGE:                func() LedgerEntry { return &Bridge{leBase: leBase{LedgerEntryType: BRIDGE}} },
	CHECK:                 func() LedgerEntry { return &Check{leBase: leBase{LedgerEntryType: CHECK}} },
	DEPOSIT_PRE_AUTH:      func() LedgerEntry { return &DepositPreAuth{leBase: leBase{LedgerEntryType: DEPOSIT_PRE_AUTH}} },
	DID:                   func() LedgerEntry { return &Did{leBase: leBase{LedgerEntryType: DID}} },
	DIRECTORY:             func() LedgerEntry { return &Directory{leBase: leBase{LedgerEntryType: DIRECTORY}} },
	ESCROW:                func() LedgerEntry { return &Escrow{leBase: leBase{LedgerEntryType: ESCROW}} },
	FEE_SETTINGS:          func() LedgerEntry { return &FeeSettings{leBase: leBase{LedgerEntryType: FEE_SETTINGS}} },
	LEDGER_HASHES:         func() LedgerEntry { return &LedgerHashes{leBase: leBase{LedgerEntryType: LEDGER_HASHES}} },
	NEGATIVE_UNL:          func() LedgerEntry { return &NegativeUNL{leBase: leBase{LedgerEntryType: NEGATIVE_UNL}} },
	NFTOKEN_OFFER:         func() LedgerEntry { return &NFTokenOffer{leBase: leBase{LedgerEntryType: NFTOKEN_OFFER}} },
	NFTOKEN_PAGE:          func() LedgerEntry { return &NFTokenPage{leBase: leBase{LedgerEntryType: NFTOKEN_PAGE}} },
	OFFER:                 func() LedgerEntry { return &Offer{leBase: leBase{LedgerEntryType: OFFER}} },
	ORACLE:                func() LedgerEntry { return &Oracle{leBase: leBase{LedgerEntryType: ORACLE}} },
	PAY_CHANNEL:           func() LedgerEntry { return &PayChannel{leBase: leBase{LedgerEntryType: PAY_CHANNEL}} },
	RIPPLE_STATE:          func() LedgerEntry { return &RippleState{leBase: leBase{LedgerEntryType: RIPPLE_STATE}} },
	SIGNER_LIST:           func() LedgerEntry { return &SignerList{leBase: leBase{LedgerEntryType: SIGNER_LIST}} },
	TICKET:                func() LedgerEntry { return &Ticket{leBase: leBase{LedgerEntryType: TICKET}} },
	XCHAIN_OWNED_CLAIM_ID: func() LedgerEntry { return &XChainOwnedClaimID{leBase: leBase{LedgerEntryType: XCHAIN_OWNED_CLAIM_ID}} },
	XCHAIN_OWNED_CREATE_ACCOUNT_CLAIM_ID: func() LedgerEntry {
		return &XChainOwnedCreateAccountClaimID{leBase: leBase{LedgerEntryType: XCHAIN_OWNED_CREATE_ACCOUNT_CLAIM_ID}}
	},
	CREDENTIAL: func() LedgerEntry { return &Credential{leBase: leBase{LedgerEntryType: CREDENTIAL}} },
}

var TxFactory = [...]func() Transaction{
	PAYMENT:              func() Transaction { return &Payment{TxBase: TxBase{TransactionType: PAYMENT}} },
	ACCOUNT_SET:          func() Transaction { return &AccountSet{TxBase: TxBase{TransactionType: ACCOUNT_SET}} },
	ACCOUNT_DELETE:       func() Transaction { return &AccountDelete{TxBase: TxBase{TransactionType: ACCOUNT_DELETE}} },
	SET_REGULAR_KEY:      func() Transaction { return &SetRegularKey{TxBase: TxBase{TransactionType: SET_REGULAR_KEY}} },
	OFFER_CREATE:         func() Transaction { return &OfferCreate{TxBase: TxBase{TransactionType: OFFER_CREATE}} },
	OFFER_CANCEL:         func() Transaction { return &OfferCancel{TxBase: TxBase{TransactionType: OFFER_CANCEL}} },
	TRUST_SET:            func() Transaction { return &TrustSet{TxBase: TxBase{TransactionType: TRUST_SET}} },
	AMENDMENT:            func() Transaction { return &Amendment{TxBase: TxBase{TransactionType: AMENDMENT}} },
	SET_FEE:              func() Transaction { return &SetFee{TxBase: TxBase{TransactionType: SET_FEE}} },
	UNL_MODIFY:           func() Transaction { return &UNLModify{TxBase: TxBase{TransactionType: UNL_MODIFY}} },
	TICKET_CREATE:        func() Transaction { return &TicketCreate{TxBase: TxBase{TransactionType: TICKET_CREATE}} },
	ESCROW_CREATE:        func() Transaction { return &EscrowCreate{TxBase: TxBase{TransactionType: ESCROW_CREATE}} },
	ESCROW_FINISH:        func() Transaction { return &EscrowFinish{TxBase: TxBase{TransactionType: ESCROW_FINISH}} },
	ESCROW_CANCEL:        func() Transaction { return &EscrowCancel{TxBase: TxBase{TransactionType: ESCROW_CANCEL}} },
	SIGNER_LIST_SET:      func() Transaction { return &SignerListSet{TxBase: TxBase{TransactionType: SIGNER_LIST_SET}} },
	PAYCHAN_CREATE:       func() Transaction { return &PaymentChannelCreate{TxBase: TxBase{TransactionType: PAYCHAN_CREATE}} },
	PAYCHAN_FUND:         func() Transaction { return &PaymentChannelFund{TxBase: TxBase{TransactionType: PAYCHAN_FUND}} },
	PAYCHAN_CLAIM:        func() Transaction { return &PaymentChannelClaim{TxBase: TxBase{TransactionType: PAYCHAN_CLAIM}} },
	CHECK_CREATE:         func() Transaction { return &CheckCreate{TxBase: TxBase{TransactionType: CHECK_CREATE}} },
	CHECK_CASH:           func() Transaction { return &CheckCash{TxBase: TxBase{TransactionType: CHECK_CASH}} },
	CHECK_CANCEL:         func() Transaction { return &CheckCancel{TxBase: TxBase{TransactionType: CHECK_CANCEL}} },
	SET_DEPOSIT_PREAUTH:  func() Transaction { return &SetDepositPreAuth{TxBase: TxBase{TransactionType: SET_DEPOSIT_PREAUTH}} },
	NFTOKEN_MINT:         func() Transaction { return &NFTokenMint{TxBase: TxBase{TransactionType: NFTOKEN_MINT}} },
	NFTOKEN_BURN:         func() Transaction { return &NFTokenBurn{TxBase: TxBase{TransactionType: NFTOKEN_BURN}} },
	NFTOKEN_MODIFY:       func() Transaction { return &NFTokenModify{TxBase: TxBase{TransactionType: NFTOKEN_MODIFY}} },
	NFTOKEN_CREATE_OFFER: func() Transaction { return &NFTokenCreateOffer{TxBase: TxBase{TransactionType: NFTOKEN_CREATE_OFFER}} },
	NFTOKEN_CANCEL_OFFER: func() Transaction { return &NFTCancelOffer{TxBase: TxBase{TransactionType: NFTOKEN_CANCEL_OFFER}} },
	NFTOKEN_ACCEPT_OFFER: func() Transaction { return &NFTAcceptOffer{TxBase: TxBase{TransactionType: NFTOKEN_ACCEPT_OFFER}} },
	CLAWBACK:             func() Transaction { return &Clawback{TxBase: TxBase{TransactionType: CLAWBACK}} },
	AMM_CREATE:           func() Transaction { return &AMMCreate{TxBase: TxBase{TransactionType: AMM_CREATE}} },
	AMM_DEPOSIT:          func() Transaction { return &AMMDeposit{TxBase: TxBase{TransactionType: AMM_DEPOSIT}} },
	AMM_WITHDRAW:         func() Transaction { return &AMMWithdraw{TxBase: TxBase{TransactionType: AMM_WITHDRAW}} },
	AMM_VOTE:             func() Transaction { return &AMMVote{TxBase: TxBase{TransactionType: AMM_VOTE}} },
	AMM_BID:              func() Transaction { return &AMMBid{TxBase: TxBase{TransactionType: AMM_BID}} },
	AMM_DELETE:           func() Transaction { return &AMMDelete{TxBase: TxBase{TransactionType: AMM_DELETE}} },
	DID_SET:              func() Transaction { return &DidSet{TxBase: TxBase{TransactionType: DID_SET}} },
	DID_DELETE:           func() Transaction { return &DidDelete{TxBase: TxBase{TransactionType: DID_DELETE}} },
	CREDENTIAL_CREATE:    func() Transaction { return &CredentialCreate{TxBase: TxBase{TransactionType: CREDENTIAL_CREATE}} },
	ORACLE_SET:           func() Transaction { return &OracleSet{TxBase: TxBase{TransactionType: ORACLE_SET}} },
	ORACLE_DELETE:        func() Transaction { return &OracleDelete{TxBase: TxBase{TransactionType: ORACLE_DELETE}} },
	XCHAIN_ACCOUNT_CREATE_COMMIT: func() Transaction {
		return &XChainAccountCreateCommit{TxBase: TxBase{TransactionType: XCHAIN_ACCOUNT_CREATE_COMMIT}}
	},
	XCHAIN_ADD_ACCOUNT_CREATE_ATTESTATION: func() Transaction {
		return &XChainAddAccountCreateAttestation{TxBase: TxBase{TransactionType: XCHAIN_ADD_ACCOUNT_CREATE_ATTESTATION}}
	},
	XCHAIN_ADD_CLAIM_ATTESTATION: func() Transaction {
		return &XChainAddClaimAttestation{TxBase: TxBase{TransactionType: XCHAIN_ADD_CLAIM_ATTESTATION}}
	},
	XCHAIN_CLAIM:  func() Transaction { return &XChainClaim{TxBase: TxBase{TransactionType: XCHAIN_CLAIM}} },
	XCHAIN_COMMIT: func() Transaction { return &XChainCommit{TxBase: TxBase{TransactionType: XCHAIN_COMMIT}} },
	XCHAIN_CREATE_BRIDGE: func() Transaction {
		return &XChainCreateBridge{TxBase: TxBase{TransactionType: XCHAIN_CREATE_BRIDGE}}
	},
	XCHAIN_CREATE_CLAIM_ID: func() Transaction {
		return &XChainCreateClaimID{TxBase: TxBase{TransactionType: XCHAIN_CREATE_CLAIM_ID}}
	},
	XCHAIN_MODIFY_BRIDGE: func() Transaction {
		return &XChainModifyBridge{TxBase: TxBase{TransactionType: XCHAIN_MODIFY_BRIDGE}}
	},
}

var ledgerEntryNames = [...]string{
	ACCOUNT_ROOT:                         "AccountRoot",
	AMENDMENTS:                           "Amendments",
	AMM_LT:                               "AMM",
	BRIDGE:                               "Bridge",
	CHECK:                                "Check",
	DEPOSIT_PRE_AUTH:                     "DepositPreauth",
	DID:                                  "DID",
	DIRECTORY:                            "DirectoryNode",
	ESCROW:                               "Escrow",
	FEE_SETTINGS:                         "FeeSettings",
	LEDGER_HASHES:                        "LedgerHashes",
	NEGATIVE_UNL:                         "NegativeUNL",
	NFTOKEN_OFFER:                        "NFTokenOffer",
	NFTOKEN_PAGE:                         "NFTokenPage",
	OFFER:                                "Offer",
	ORACLE:                               "Oracle",
	PAY_CHANNEL:                          "PayChannel",
	RIPPLE_STATE:                         "RippleState",
	SIGNER_LIST:                          "SignerList",
	TICKET:                               "Ticket",
	XCHAIN_OWNED_CLAIM_ID:                "XChainOwnedClaimID",
	XCHAIN_OWNED_CREATE_ACCOUNT_CLAIM_ID: "XChainOwnedCreateAccountClaimID",
	CREDENTIAL:                           "Credential",
}

var ledgerEntryTypes = map[string]LedgerEntryType{
	"AccountRoot":                     ACCOUNT_ROOT,
	"Amendments":                      AMENDMENTS,
	"AMM":                             AMM_LT,
	"Bridge":                          BRIDGE,
	"Check":                           CHECK,
	"DID":                             DID,
	"DepositPreauth":                  DEPOSIT_PRE_AUTH,
	"DirectoryNode":                   DIRECTORY,
	"Escrow":                          ESCROW,
	"FeeSettings":                     FEE_SETTINGS,
	"LedgerHashes":                    LEDGER_HASHES,
	"NegativeUNL":                     NEGATIVE_UNL,
	"NFTokenOffer":                    NFTOKEN_OFFER,
	"NFTokenPage":                     NFTOKEN_PAGE,
	"Offer":                           OFFER,
	"Oracle":                          ORACLE,
	"PayChannel":                      PAY_CHANNEL,
	"RippleState":                     RIPPLE_STATE,
	"SignerList":                      SIGNER_LIST,
	"Ticket":                          TICKET,
	"XChainOwnedClaimID":              XCHAIN_OWNED_CLAIM_ID,
	"XChainOwnedCreateAccountClaimID": XCHAIN_OWNED_CREATE_ACCOUNT_CLAIM_ID,
	"Credential":                      CREDENTIAL,
}

var txNames = [...]string{
	PAYMENT:                               "Payment",
	ACCOUNT_SET:                           "AccountSet",
	ACCOUNT_DELETE:                        "AccountDelete",
	SET_REGULAR_KEY:                       "SetRegularKey",
	OFFER_CREATE:                          "OfferCreate",
	OFFER_CANCEL:                          "OfferCancel",
	TRUST_SET:                             "TrustSet",
	AMENDMENT:                             "EnableAmendment",
	SET_FEE:                               "SetFee",
	UNL_MODIFY:                            "UNLModify",
	TICKET_CREATE:                         "TicketCreate",
	ESCROW_CREATE:                         "EscrowCreate",
	ESCROW_FINISH:                         "EscrowFinish",
	ESCROW_CANCEL:                         "EscrowCancel",
	SIGNER_LIST_SET:                       "SignerListSet",
	PAYCHAN_CREATE:                        "PaymentChannelCreate",
	PAYCHAN_FUND:                          "PaymentChannelFund",
	PAYCHAN_CLAIM:                         "PaymentChannelClaim",
	CHECK_CREATE:                          "CheckCreate",
	CHECK_CASH:                            "CheckCash",
	CHECK_CANCEL:                          "CheckCancel",
	SET_DEPOSIT_PREAUTH:                   "DepositPreauth",
	NFTOKEN_MINT:                          "NFTokenMint",
	NFTOKEN_BURN:                          "NFTokenBurn",
	NFTOKEN_MODIFY:                        "NFTokenModify",
	NFTOKEN_CREATE_OFFER:                  "NFTokenCreateOffer",
	NFTOKEN_CANCEL_OFFER:                  "NFTokenCancelOffer",
	NFTOKEN_ACCEPT_OFFER:                  "NFTokenAcceptOffer",
	CLAWBACK:                              "Clawback",
	AMM_CREATE:                            "AMMCreate",
	AMM_DEPOSIT:                           "AMMDeposit",
	AMM_WITHDRAW:                          "AMMWithdraw",
	AMM_VOTE:                              "AMMVote",
	AMM_BID:                               "AMMBid",
	AMM_DELETE:                            "AMMDelete",
	DID_SET:                               "DIDSet",
	DID_DELETE:                            "DIDDelete",
	ORACLE_SET:                            "OracleSet",
	ORACLE_DELETE:                         "OracleDelete",
	XCHAIN_ACCOUNT_CREATE_COMMIT:          "XChainAccountCreateCommit",
	XCHAIN_ADD_ACCOUNT_CREATE_ATTESTATION: "XChainAddAccountCreateAttestation",
	XCHAIN_ADD_CLAIM_ATTESTATION:          "XChainAddClaimAttestation",
	XCHAIN_CLAIM:                          "XChainClaim",
	XCHAIN_COMMIT:                         "XChainCommit",
	XCHAIN_CREATE_BRIDGE:                  "XChainCreateBridge",
	XCHAIN_CREATE_CLAIM_ID:                "XChainCreateClaimID",
	XCHAIN_MODIFY_BRIDGE:                  "XChainModifyBridge",
	CREDENTIAL_CREATE:                     "CredentialCreate",
}

var txTypes = map[string]TransactionType{
	"Payment":                           PAYMENT,
	"AccountSet":                        ACCOUNT_SET,
	"AccountDelete":                     ACCOUNT_DELETE,
	"SetRegularKey":                     SET_REGULAR_KEY,
	"OfferCreate":                       OFFER_CREATE,
	"OfferCancel":                       OFFER_CANCEL,
	"TrustSet":                          TRUST_SET,
	"EnableAmendment":                   AMENDMENT,
	"SetFee":                            SET_FEE,
	"UNLModify":                         UNL_MODIFY,
	"TicketCreate":                      TICKET_CREATE,
	"EscrowCreate":                      ESCROW_CREATE,
	"EscrowFinish":                      ESCROW_FINISH,
	"EscrowCancel":                      ESCROW_CANCEL,
	"SignerListSet":                     SIGNER_LIST_SET,
	"PaymentChannelCreate":              PAYCHAN_CREATE,
	"PaymentChannelFund":                PAYCHAN_FUND,
	"PaymentChannelClaim":               PAYCHAN_CLAIM,
	"CheckCreate":                       CHECK_CREATE,
	"CheckCash":                         CHECK_CASH,
	"CheckCancel":                       CHECK_CANCEL,
	"DepositPreauth":                    SET_DEPOSIT_PREAUTH,
	"NFTokenMint":                       NFTOKEN_MINT,
	"NFTokenBurn":                       NFTOKEN_BURN,
	"NFTokenModify":                     NFTOKEN_MODIFY,
	"NFTokenCreateOffer":                NFTOKEN_CREATE_OFFER,
	"NFTokenCancelOffer":                NFTOKEN_CANCEL_OFFER,
	"NFTokenAcceptOffer":                NFTOKEN_ACCEPT_OFFER,
	"Clawback":                          CLAWBACK,
	"AMMCreate":                         AMM_CREATE,
	"AMMDeposit":                        AMM_DEPOSIT,
	"AMMWithdraw":                       AMM_WITHDRAW,
	"AMMVote":                           AMM_VOTE,
	"AMMBid":                            AMM_BID,
	"AMMDelete":                         AMM_DELETE,
	"DIDSet":                            DID_SET,
	"DIDDelete":                         DID_DELETE,
	"OracleSet":                         ORACLE_SET,
	"OracleDelete":                      ORACLE_DELETE,
	"XChainAccountCreateCommit":         XCHAIN_ACCOUNT_CREATE_COMMIT,
	"XChainAddAccountCreateAttestation": XCHAIN_ADD_ACCOUNT_CREATE_ATTESTATION,
	"XChainAddClaimAttestation":         XCHAIN_ADD_CLAIM_ATTESTATION,
	"XChainClaim":                       XCHAIN_CLAIM,
	"XChainCommit":                      XCHAIN_COMMIT,
	"XChainCreateBridge":                XCHAIN_CREATE_BRIDGE,
	"XChainCreateClaimID":               XCHAIN_CREATE_CLAIM_ID,
	"XChainModifyBridge":                XCHAIN_MODIFY_BRIDGE,
	"CredentialCreate":                  CREDENTIAL_CREATE,
}

var HashableTypes []string

func init() {
	HashableTypes = append(HashableTypes, NT_TRANSACTION_NODE.String())
	for _, typ := range txNames {
		if len(typ) > 0 {
			HashableTypes = append(HashableTypes, typ)
		}
	}
	HashableTypes = append(HashableTypes, NT_ACCOUNT_NODE.String())
	for _, typ := range ledgerEntryNames {
		if len(typ) > 0 {
			HashableTypes = append(HashableTypes, typ)
		}
	}
}

func (t TransactionType) String() string {
	return txNames[t]
}

func (le LedgerEntryType) String() string {
	return ledgerEntryNames[le]
}

func GetTxFactoryByType(txType string) func() Transaction {
	return TxFactory[txTypes[txType]]
}

func GetLedgerEntryFactoryByType(leType string) func() LedgerEntry {
	return LedgerEntryFactory[ledgerEntryTypes[leType]]
}
