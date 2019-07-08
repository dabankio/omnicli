package btccli

// 固定常量
const (
	RPCUser   = "rpcusr"
	RPCPasswd = "233"

	RPCPortRegtest = 18443

	OmniBinPathEnv = "OMNI_BIN_PATH"

	CmdParamRegtest = "-regtest"
)

var (
	// PresetAccounts 一些预置账号
	PresetAccounts = []Addr{
		{Address: "mt6WnfniPEKziTgmAdebBkT33vAcM7jJFC", Privkey: "cTBichWb6XBPa6F6NUc3gkLjYDufRWx74CnJWEypxuFNsXd6dJNb", Pubkey: "032affe4397f911be5d4ff84531d84b7d9498263731f8c205c1082e73abcc60955"},
		{Address: "miP2AS7DQmQuc9WSYzz5uJrFh6xhe4TvYw", Privkey: "cVWuAHG981YURgJCFRYHuVdApgKtmciujVmkYFkZnaVeMLXz3pKX", Pubkey: "0264f183744e1dd7523ead4eef81b99e13870a8234724ee9035aee0ccfa70f1b1a"},
		{Address: "n4ekdCUgRUCZxAiVTiCKfKVU4fz5SewcKr", Privkey: "cNczugkPEDYcCm1UhukMJp1XCGdtqez4cVfzK4yBfcxtYbF2Ui7r", Pubkey: "0383b83e84538b334df3f8820539053e635e30ba05b85a05850302971d9b3ece3c"},
		{Address: "mrmsAWtDKgqGEppz2ULnd1yjthinGgZmY9", Privkey: "cNkMpgrsvvRjzkNYvBmMBWLbusFgCWJNrbfzhecmmSaKN7q9cgwx", Pubkey: "03ca007eefc39e3406d2131cf498efd5cca7fe3b3c5b61625b58a8342e84a4f947"},
		{Address: "mtxrtZn5RnMGXADrfUaMcyrb8EJ1c7276W", Privkey: "cN4pHDRjKbCjHot5iyQorPRVQWQqfRuWTFbCgPCz2UG2y1K9oHv9", Pubkey: "02f0744be0f11fd0db10bd669bb20f29e8fdbfdc39b0b9d1f8c81be2f42d7068a7"},
		{Address: "n162u2V7Sav6mUC1baJ1u7hGrBtsSLdszq", Privkey: "cN8Sq2fjq5bnUsRTQNL6CSdwxA7d2FmSsSG9r8h1DBwfY7eaz5EA", Pubkey: "02e5dcd14da054866160fa9adb1af3e04c21eae1d2677b46bded4d9b7f7b5a22b3"},
		{Address: "mopaTypypdaFh3u5nV16vsqZb22DcqM1pf", Privkey: "cUCXumLXtHHjLBr3mn888uyBvx45ELwznnzPzuFCvk9SThUi7jUJ", Pubkey: "0285c13d0769efbf118a0b7f5b87150f8f87fbe07a7acca8fe3a92e8b6541b0a65"},
		{Address: "mxFFWmHfp1Gz657QZeVuGiTCeEk71Z6aYX", Privkey: "cRsF4QaXXVwBNgunPEzNgEHkxz8GoeWzZ8jrk8JM46unBZaaiXmP", Pubkey: "03c935d1934055b83e1b19e60d16a8c9d4f89392aa3057d32b7d88a5ea36b098f6"},
		{Address: "mqbyybrXRZwHUET7taiShJ1S6dScnBd7yg", Privkey: "cV7Z95EMA7G5vrQ9bCe1tZRmL5eMncimvVXPjY9FuR7UyKzWvdZH", Pubkey: "0326b988f6da309846aded94a56f3230e1bdd1494878343830e89053ddd596e069"},
		{Address: "mzfveH16kQ7rNdPHNJMrkpWqjo7hVdpVHC", Privkey: "cPEnbKH87636VhGhsVQxHBV6RmPyp3d1NL9q9XQJGXpyMhyUJoMJ", Pubkey: "026c27bbb52665e6d1d3d1f5ccb8ac7b7b27ee8636019af14272ebaa0308b800b2"},
	}
)
