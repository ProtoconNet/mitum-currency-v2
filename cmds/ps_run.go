package cmds

import (
	"github.com/spikeekips/mitum/launch"
	"github.com/spikeekips/mitum/util/ps"
)

var (
	PNameDigest           = ps.Name("digest")
	PNameDigestStart      = ps.Name("digest_star")
	PNameMongoDBsDataBase = ps.Name("mongodb_database")
	PNameDigestDataBase   = ps.Name("digest_database")
)

func DefaultRunPS() *ps.PS {
	pps := ps.NewPS("cmd-run")

	_ = pps.
		AddOK(launch.PNameEncoder, PEncoder, nil).
		AddOK(launch.PNameDesign, PLoadDesign, nil, launch.PNameEncoder).
		AddOK(launch.PNameTimeSyncer, PStartTimeSyncer /*launch.PCloseTimeSyncer, */, nil, launch.PNameDesign).
		AddOK(launch.PNameLocal, PLocal, nil, launch.PNameDesign).
		AddOK(launch.PNameStorage, launch.PStorage, nil, launch.PNameLocal).
		AddOK(launch.PNameProposalMaker, launch.PProposalMaker, nil, launch.PNameStorage).
		AddOK(launch.PNameNetwork, PNetwork, nil, launch.PNameStorage).
		AddOK(launch.PNameMemberlist, PMemberlist, nil, launch.PNameNetwork).
		AddOK(launch.PNameStartNetwork, launch.PStartNetwork, launch.PCloseNetwork, launch.PNameStates).
		AddOK(launch.PNameStartStorage, launch.PStartStorage, launch.PCloseStorage, launch.PNameStartNetwork).
		AddOK(launch.PNameStartMemberlist, launch.PStartMemberlist, launch.PCloseMemberlist, launch.PNameStartNetwork).
		AddOK(launch.PNameStartSyncSourceChecker, launch.PStartSyncSourceChecker, launch.PCloseSyncSourceChecker, launch.PNameStartNetwork).
		AddOK(launch.PNameStartLastConsensusNodesWatcher,
			launch.PStartLastConsensusNodesWatcher, launch.PCloseLastConsensusNodesWatcher, launch.PNameStartNetwork).
		AddOK(launch.PNameStates, launch.PStates, nil, launch.PNameNetwork).
		AddOK(launch.PNameStatesReady, nil, launch.PCloseStates,
			launch.PNameStartStorage,
			launch.PNameStartSyncSourceChecker,
			launch.PNameStartLastConsensusNodesWatcher,
			launch.PNameStartMemberlist,
			launch.PNameStartNetwork,
			launch.PNameStates).
		AddOK(PNameMongoDBsDataBase, ProcessDatabase, nil, launch.PNameDesign, launch.PNameStorage).
		AddOK(PNameDigestDataBase, ProcessDigestDatabase, nil, PNameMongoDBsDataBase).
		AddOK(PNameDigester, ProcessDigester, nil, PNameDigestDataBase).
		AddOK(PNameDigest, ProcessDigestAPI, nil, launch.PNameDesign, PNameDigestDataBase, launch.PNameMemberlist).
		AddOK(PNameDigestStart, ProcessStartDigestAPI, nil, PNameDigest).
		AddOK(PNameStartDigester, ProcessStartDigester, nil, PNameDigestStart)

	_ = pps.POK(launch.PNameEncoder).
		PostAddOK(launch.PNameAddHinters, PAddHinters)

	_ = pps.POK(launch.PNameDesign).
		PostAddOK(launch.PNameCheckDesign, PCheckDesign)

	_ = pps.POK(launch.PNameLocal).
		PostAddOK(launch.PNameDiscoveryFlag, launch.PDiscoveryFlag)

	_ = pps.POK(launch.PNameStorage).
		PreAddOK(launch.PNameCheckLocalFS, PCheckLocalFS).
		PreAddOK(launch.PNameLoadDatabase, PLoadDatabase).
		PostAddOK(launch.PNameCheckLeveldbStorage, launch.PCheckLeveldbStorage).
		PostAddOK(launch.PNameCheckLoadFromDatabase, PLoadFromDatabase).
		PostAddOK(launch.PNameGetSuffrageFromDatabaseeFunc, launch.PGetSuffrageFromDatabaseFunc).
		PostAddOK(launch.PNameNodeInfo, PNodeInfo)

	_ = pps.POK(launch.PNameNetwork).
		PreAddOK(launch.PNameQuicstreamClient, launch.PQuicstreamClient).
		PostAddOK(launch.PNameSyncSourceChecker, PSyncSourceChecker).
		PostAddOK(launch.PNameSuffrageCandidateLimiterSet, launch.PSuffrageCandidateLimiterSet)

	_ = pps.POK(launch.PNameMemberlist).
		PreAddOK(launch.PNameLastConsensusNodesWatcher, launch.PLastConsensusNodesWatcher).
		PostAddOK(launch.PNameBallotbox, launch.PBallotbox).
		PostAddOK(launch.PNameLongRunningMemberlistJoin, launch.PLongRunningMemberlistJoin).
		PostAddOK(launch.PNameCallbackBroadcaster, PCallbackBroadcaster).
		PostAddOK(launch.PNameSuffrageVoting, launch.PSuffrageVoting)

	_ = pps.POK(launch.PNameStates).
		PreAddOK(launch.PNameProposerSelector, launch.PProposerSelector).
		PreAddOK(launch.PNameOperationProcessorsMap, POperationProcessorsMap).
		PreAddOK(launch.PNameNetworkHandlers, PNetworkHandlers).
		PreAddOK(launch.PNameNodeInConsensusNodesFunc, launch.PNodeInConsensusNodesFunc).
		PreAddOK(launch.PNameProposalProcessors, PProposalProcessors).
		PostAddOK(launch.PNamePatchLastConsensusNodesWatcher, launch.PPatchLastConsensusNodesWatcher).
		PostAddOK(launch.PNameStatesSetHandlers, PStatesSetHandlers).
		PostAddOK(launch.PNameWatchDesign, PWatchDesign).
		PostAddOK(launch.PNamePatchMemberlist, launch.PPatchMemberlist)

	return pps
}
