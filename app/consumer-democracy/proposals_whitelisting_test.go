package app_test

//COMMENTED OUT BECAUSE OF SOFT OPT OUT FEATURE
//func TestDemocracyGovernanceWhitelistingKeys(t *testing.T) {
//	chain := ibctesting.NewTestChain(t, ibctesting.NewCoordinator(t, 0),
//		icstestingutils.DemocracyConsumerAppIniter, "test")
//	paramKeeper := chain.App.(*appConsumer.App).ParamsKeeper
//	for paramKey := range appConsumer.WhitelistedParams {
//		ss, ok := paramKeeper.GetSubspace(paramKey.Subspace)
//		require.True(t, ok, "Unknown subspace %s", paramKey.Subspace)
//		hasKey := ss.Has(chain.GetContext(), []byte(paramKey.Key))
//		require.True(t, hasKey, "Invalid key %s for subspace %s", paramKey.Key, paramKey.Subspace)
//	}
//}
