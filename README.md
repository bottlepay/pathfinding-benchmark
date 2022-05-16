# Pathfinding Benchmark

The goal of the code in this repository is to provide a node-implementation
independent benchmark for pathfinding algorithms.

Relevant files:

* [graph.yml](graph.yml) defines the test graph. The node `start` is special
  because this is where payments are made from and for which the implementation
  can be switched between CLN and LND. All other nodes in the test graph are LND
  based.

  In this file, there's also a section `tests` that defines the test payments.

* [cmd/gencluster](cmd/gencluster) contains code to generate the
  [docker-compose](docker-compose.yml) file from `graph.yml`.

* [cmd/testrunner/cmd_run.go](cmd/testrunner/cmd_run.go) is the automated test
  that sets up channels and executes the payments.

* [run.sh cln | lnd](run.sh) fires it all up. The argument controls the
  implementation that is used for the `start` node.

  If all goes well, output should be similar to:

```
testrunner_1                | Starting test.
testrunner_1                | 2022-05-16T08:12:11.194Z	INFO	Attempting to connect to bitcoind
testrunner_1                | 2022-05-16T08:12:11.207Z	INFO	Connected to bitcoind
testrunner_1                | 2022-05-16T08:12:11.207Z	INFO	Creating bitcoind wallet
testrunner_1                | 2022-05-16T08:12:21.378Z	INFO	Bitcoin address	{"address": "bcrt1qc4rrvj5nmta6jrpff5tuk40l08fp7r065yd7j4"}
testrunner_1                | 2022-05-16T08:12:21.378Z	INFO	Activate segwit
testrunner_1                | 2022-05-16T08:12:29.867Z	INFO	Fund senders
testrunner_1                | 2022-05-16T08:12:29.867Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_3"}
testrunner_1                | 2022-05-16T08:12:33.064Z	INFO	Connected to lnd	{"node": "no_liquidity_3", "key": "034c887a1d347eab4e0c98d1a308c2c3a2a0dc46f200e7c9b665d808d9a8d0ace0"}
testrunner_1                | 2022-05-16T08:12:33.075Z	INFO	Generated funding address	{"node": "no_liquidity_3", "address": "bcrt1qy3caf2fw56rch5h3c2nrj8n6620eqnmcjkanuj"}
testrunner_1                | 2022-05-16T08:12:34.672Z	INFO	Attempting to connect to lnd	{"node": "node1_1"}
testrunner_1                | 2022-05-16T08:12:34.723Z	INFO	Connected to lnd	{"node": "node1_1", "key": "03071ce2ce4e66ce7743c80d97a1a98cd94ed5c8a4d8a5aa6c8d691a3a4369c23e"}
testrunner_1                | 2022-05-16T08:12:34.729Z	INFO	Generated funding address	{"node": "node1_1", "address": "bcrt1q6qwdf89h9n7wgl6p3xpfs3x9lpsaupremz2jtu"}
testrunner_1                | 2022-05-16T08:12:36.297Z	INFO	Attempting to connect to lnd	{"node": "node3_1"}
testrunner_1                | 2022-05-16T08:12:36.325Z	INFO	Connected to lnd	{"node": "node3_1", "key": "027deb0a8ff6a857a51b8c34b1826ef32deb85b2b8843ab83cd7fb8aa8ee84792e"}
testrunner_1                | 2022-05-16T08:12:36.330Z	INFO	Generated funding address	{"node": "node3_1", "address": "bcrt1qnfgc66xmzwda3cf6r0lq78mk4nka3q687rnk7c"}
testrunner_1                | 2022-05-16T08:12:37.485Z	INFO	Attempting to connect to lnd	{"node": "node4_0"}
testrunner_1                | 2022-05-16T08:12:37.497Z	INFO	Connected to lnd	{"node": "node4_0", "key": "0325c9adb14fd2052a0705a8298316fbff71afe3a577d6fcf413c47160596d0b14"}
testrunner_1                | 2022-05-16T08:12:37.499Z	INFO	Generated funding address	{"node": "node4_0", "address": "bcrt1qmqqmmnzsxfelhh68v5tqagdcjp0le8a92s3sgr"}
testrunner_1                | 2022-05-16T08:12:38.498Z	INFO	Attempting to connect to lnd	{"node": "node4_1"}
testrunner_1                | 2022-05-16T08:12:38.509Z	INFO	Connected to lnd	{"node": "node4_1", "key": "02da9062200f01617131d75bf45775340323aa31dbb992255fa5513d90f8857787"}
testrunner_1                | 2022-05-16T08:12:38.512Z	INFO	Generated funding address	{"node": "node4_1", "address": "bcrt1qg8xhfqs8ddksl0c4qewl2aj0x3gs5ku8z3vqlq"}
testrunner_1                | 2022-05-16T08:12:39.489Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_2"}
testrunner_1                | 2022-05-16T08:12:39.501Z	INFO	Connected to lnd	{"node": "no_liquidity_2", "key": "0374e63ee63dcabe31b8875d8d8a0792d2b38918ca5014502f9585a51738e42c58"}
testrunner_1                | 2022-05-16T08:12:39.504Z	INFO	Generated funding address	{"node": "no_liquidity_2", "address": "bcrt1q09z534ectcztn4rh6y7svsv757y4vhvky6tapl"}
testrunner_1                | 2022-05-16T08:12:40.485Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_4"}
testrunner_1                | 2022-05-16T08:12:40.497Z	INFO	Connected to lnd	{"node": "no_liquidity_4", "key": "024d788fb7f741bf9e8d353819fda5a715284ed6882a5a2485b0463b8c605e0f0e"}
testrunner_1                | 2022-05-16T08:12:40.501Z	INFO	Generated funding address	{"node": "no_liquidity_4", "address": "bcrt1qtz6wf7sx94vrl4klzjkwa459tmy8ytdtna6ax4"}
testrunner_1                | 2022-05-16T08:12:41.538Z	INFO	Attempting to connect to lnd	{"node": "destination1"}
testrunner_1                | 2022-05-16T08:12:41.549Z	INFO	Connected to lnd	{"node": "destination1", "key": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b"}
testrunner_1                | 2022-05-16T08:12:41.553Z	INFO	Generated funding address	{"node": "destination1", "address": "bcrt1qgrf7p5x250d5mvvsw25s5rmdhstkyu5sfchur9"}
testrunner_1                | 2022-05-16T08:12:42.525Z	INFO	Attempting to connect to lnd	{"node": "start"}
testrunner_1                | 2022-05-16T08:12:42.535Z	INFO	Connected to lnd	{"node": "start", "key": "023101573b03fcbb436ddb197a68764f8c13020e89ff26c821213885fb7a2ae844"}
testrunner_1                | 2022-05-16T08:12:42.538Z	INFO	Generated funding address	{"node": "start", "address": "bcrt1qt29seaefwf4kas6e83h3t6cnqv6sww2lznd2xs"}
testrunner_1                | 2022-05-16T08:12:43.506Z	INFO	Attempting to connect to lnd	{"node": "node2_0"}
testrunner_1                | 2022-05-16T08:12:43.517Z	INFO	Connected to lnd	{"node": "node2_0", "key": "02b1a1f35e76116a0460e100f65b952c4ebe6e04c973a610c36f7950a5f91bdd64"}
testrunner_1                | 2022-05-16T08:12:43.520Z	INFO	Generated funding address	{"node": "node2_0", "address": "bcrt1q6zppc04ratvv8twu9eexfchmhy54xrvzfnz6xs"}
testrunner_1                | 2022-05-16T08:12:44.484Z	INFO	Attempting to connect to lnd	{"node": "node3_0"}
testrunner_1                | 2022-05-16T08:12:44.496Z	INFO	Connected to lnd	{"node": "node3_0", "key": "031ce4fd8535b25aa8c9090c21c282bfce10e3143c5573553c18f4e3fb14d2d1b2"}
testrunner_1                | 2022-05-16T08:12:44.499Z	INFO	Generated funding address	{"node": "node3_0", "address": "bcrt1q2mtxm6erx9prtaresnf59u4ykq86kmp4m3vnl6"}
testrunner_1                | 2022-05-16T08:12:45.490Z	INFO	Attempting to connect to lnd	{"node": "black_hole_indirect"}
testrunner_1                | 2022-05-16T08:12:45.501Z	INFO	Connected to lnd	{"node": "black_hole_indirect", "key": "02021eb03f5cb8dd028a63e41388acc89f6fece860794aaa3e8ad885acf040c13c"}
testrunner_1                | 2022-05-16T08:12:45.504Z	INFO	Generated funding address	{"node": "black_hole_indirect", "address": "bcrt1qfsg2yntj0qta0lghum627mqcwutvt5x8efnvej"}
testrunner_1                | 2022-05-16T08:12:46.474Z	INFO	Attempting to connect to lnd	{"node": "destination2"}
testrunner_1                | 2022-05-16T08:12:46.484Z	INFO	Connected to lnd	{"node": "destination2", "key": "027e4933bc5b4de4f22635c458f4735bb447933a8d6df0af49292e2c69d9bfdead"}
testrunner_1                | 2022-05-16T08:12:46.487Z	INFO	Generated funding address	{"node": "destination2", "address": "bcrt1qpst7tc84qykd4zd6puaagm7hepg83xaz2x362w"}
testrunner_1                | 2022-05-16T08:12:47.490Z	INFO	Attempting to connect to lnd	{"node": "node0_1"}
testrunner_1                | 2022-05-16T08:12:47.501Z	INFO	Connected to lnd	{"node": "node0_1", "key": "036fa23150634644b42442bf375c46e845083cb33ebd9e430dfe8ab877de175296"}
testrunner_1                | 2022-05-16T08:12:47.504Z	INFO	Generated funding address	{"node": "node0_1", "address": "bcrt1q99hj0hypf2qazeqfr8t96wnfawfmf3dx48knz5"}
testrunner_1                | 2022-05-16T08:12:48.501Z	INFO	Attempting to connect to lnd	{"node": "node1_0"}
testrunner_1                | 2022-05-16T08:12:48.512Z	INFO	Connected to lnd	{"node": "node1_0", "key": "03b0e27c266d3982eca208fc435fb6c0109acc6e9c023dc4664470ffd01da1e3e2"}
testrunner_1                | 2022-05-16T08:12:48.515Z	INFO	Generated funding address	{"node": "node1_0", "address": "bcrt1q038cdhc7fc90tesrpukdppjvf5em6h4ay7g6rn"}
testrunner_1                | 2022-05-16T08:12:49.493Z	INFO	Attempting to connect to lnd	{"node": "black_hole"}
testrunner_1                | 2022-05-16T08:12:49.504Z	INFO	Connected to lnd	{"node": "black_hole", "key": "0256bd7a7cd87c5179bcff888704e75ee7e620e042d5d16e861600516c5dd4b71c"}
testrunner_1                | 2022-05-16T08:12:49.508Z	INFO	Generated funding address	{"node": "black_hole", "address": "bcrt1qlu5d9suxem68z0pthv5vnr6kyfwjn4lqr0d0z7"}
testrunner_1                | 2022-05-16T08:12:50.524Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_0"}
testrunner_1                | 2022-05-16T08:12:50.540Z	INFO	Connected to lnd	{"node": "no_liquidity_0", "key": "030e5e975e438d9c65f663d8bce24f9352466e8e9376727330369da0451f10cfe8"}
testrunner_1                | 2022-05-16T08:12:50.544Z	INFO	Generated funding address	{"node": "no_liquidity_0", "address": "bcrt1qppv0djqrptl7dlnjr6dugf5c8uglwx9qaccsln"}
testrunner_1                | 2022-05-16T08:12:51.552Z	INFO	Attempting to connect to lnd	{"node": "node0_0"}
testrunner_1                | 2022-05-16T08:12:51.564Z	INFO	Connected to lnd	{"node": "node0_0", "key": "03ce56a2e81b6491420b0063eb11a7e13a7af32fbcb0a348ce4f1d578ad1e25064"}
testrunner_1                | 2022-05-16T08:12:51.566Z	INFO	Generated funding address	{"node": "node0_0", "address": "bcrt1q52gsqk2h8ux4ae9pf4wlhpg50hmcw3pf67hlrd"}
testrunner_1                | 2022-05-16T08:12:52.551Z	INFO	Attempting to connect to lnd	{"node": "node2_1"}
testrunner_1                | 2022-05-16T08:12:52.562Z	INFO	Connected to lnd	{"node": "node2_1", "key": "020d9e75637de128a357844440fa37ff7adb99df4dc89dddd85dfc29dcd2d95c32"}
testrunner_1                | 2022-05-16T08:12:52.565Z	INFO	Generated funding address	{"node": "node2_1", "address": "bcrt1qc0huugll4hmaa3hcadzkm6t9s3c4lzmjzs08ds"}
testrunner_1                | 2022-05-16T08:12:53.544Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_1"}
testrunner_1                | 2022-05-16T08:12:53.556Z	INFO	Connected to lnd	{"node": "no_liquidity_1", "key": "0225087786087a9db4a8b32fd35761babacf7bb8d13732dace57a132ea907c3950"}
testrunner_1                | 2022-05-16T08:12:53.559Z	INFO	Generated funding address	{"node": "no_liquidity_1", "address": "bcrt1qczy27kmqrau68lztgjt6aqx2xqpxmp3sr3s2gl"}
testrunner_1                | 2022-05-16T08:12:54.555Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_5"}
testrunner_1                | 2022-05-16T08:12:54.567Z	INFO	Connected to lnd	{"node": "no_liquidity_5", "key": "03db61e38801e90797034d9b2488ec3c6ac7372d85293270f2bdbbdfb085988a6e"}
testrunner_1                | 2022-05-16T08:12:54.570Z	INFO	Generated funding address	{"node": "no_liquidity_5", "address": "bcrt1qk066jkhpqntzswghn0mqaff2njsg2fzf4xaaua"}
testrunner_1                | 2022-05-16T08:12:55.589Z	INFO	Wait for coin and open channels
testrunner_1                | 2022-05-16T08:12:56.686Z	INFO	Connecting	{"node": "no_liquidity_3", "peer": "destination1", "peerPubKey": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b", "host": "node_destination1:9735"}
testrunner_1                | 2022-05-16T08:12:56.691Z	INFO	Open channel	{"node": "no_liquidity_3", "peer": "destination1"}
testrunner_1                | 2022-05-16T08:12:56.766Z	INFO	Connecting	{"node": "no_liquidity_3", "peer": "destination2", "peerPubKey": "027e4933bc5b4de4f22635c458f4735bb447933a8d6df0af49292e2c69d9bfdead", "host": "node_destination2:9735"}
testrunner_1                | 2022-05-16T08:12:56.771Z	INFO	Open channel	{"node": "no_liquidity_3", "peer": "destination2"}
testrunner_1                | 2022-05-16T08:12:56.923Z	INFO	Connecting	{"node": "node1_1", "peer": "node2_0", "peerPubKey": "02b1a1f35e76116a0460e100f65b952c4ebe6e04c973a610c36f7950a5f91bdd64", "host": "node_node2_0:9735"}
testrunner_1                | 2022-05-16T08:12:56.929Z	INFO	Open channel	{"node": "node1_1", "peer": "node2_0"}
testrunner_1                | 2022-05-16T08:12:57.000Z	INFO	Connecting	{"node": "node1_1", "peer": "node2_1", "peerPubKey": "020d9e75637de128a357844440fa37ff7adb99df4dc89dddd85dfc29dcd2d95c32", "host": "node_node2_1:9735"}
testrunner_1                | 2022-05-16T08:12:57.007Z	INFO	Open channel	{"node": "node1_1", "peer": "node2_1"}
testrunner_1                | 2022-05-16T08:12:57.158Z	INFO	Connecting	{"node": "node3_1", "peer": "node4_0", "peerPubKey": "0325c9adb14fd2052a0705a8298316fbff71afe3a577d6fcf413c47160596d0b14", "host": "node_node4_0:9735"}
testrunner_1                | 2022-05-16T08:12:57.163Z	INFO	Open channel	{"node": "node3_1", "peer": "node4_0"}
testrunner_1                | 2022-05-16T08:12:57.240Z	INFO	Connecting	{"node": "node3_1", "peer": "node4_1", "peerPubKey": "02da9062200f01617131d75bf45775340323aa31dbb992255fa5513d90f8857787", "host": "node_node4_1:9735"}
testrunner_1                | 2022-05-16T08:12:57.247Z	INFO	Open channel	{"node": "node3_1", "peer": "node4_1"}
testrunner_1                | 2022-05-16T08:12:57.319Z	INFO	Connecting	{"node": "node3_1", "peer": "destination2", "peerPubKey": "027e4933bc5b4de4f22635c458f4735bb447933a8d6df0af49292e2c69d9bfdead", "host": "node_destination2:9735"}
testrunner_1                | 2022-05-16T08:12:57.324Z	INFO	Open channel	{"node": "node3_1", "peer": "destination2"}
testrunner_1                | 2022-05-16T08:12:57.466Z	INFO	Connecting	{"node": "node4_0", "peer": "destination1", "peerPubKey": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b", "host": "node_destination1:9735"}
testrunner_1                | 2022-05-16T08:12:57.471Z	INFO	Open channel	{"node": "node4_0", "peer": "destination1"}
testrunner_1                | 2022-05-16T08:12:57.615Z	INFO	Connecting	{"node": "node4_1", "peer": "destination1", "peerPubKey": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b", "host": "node_destination1:9735"}
testrunner_1                | 2022-05-16T08:12:57.656Z	INFO	Open channel	{"node": "node4_1", "peer": "destination1"}
testrunner_1                | 2022-05-16T08:12:57.804Z	INFO	Connecting	{"node": "no_liquidity_2", "peer": "destination1", "peerPubKey": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b", "host": "node_destination1:9735"}
testrunner_1                | 2022-05-16T08:12:57.810Z	INFO	Open channel	{"node": "no_liquidity_2", "peer": "destination1"}
testrunner_1                | 2022-05-16T08:12:57.883Z	INFO	Connecting	{"node": "no_liquidity_2", "peer": "destination2", "peerPubKey": "027e4933bc5b4de4f22635c458f4735bb447933a8d6df0af49292e2c69d9bfdead", "host": "node_destination2:9735"}
testrunner_1                | 2022-05-16T08:12:57.891Z	INFO	Open channel	{"node": "no_liquidity_2", "peer": "destination2"}
testrunner_1                | 2022-05-16T08:12:58.040Z	INFO	Connecting	{"node": "no_liquidity_4", "peer": "destination1", "peerPubKey": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b", "host": "node_destination1:9735"}
testrunner_1                | 2022-05-16T08:12:58.045Z	INFO	Open channel	{"node": "no_liquidity_4", "peer": "destination1"}
testrunner_1                | 2022-05-16T08:12:58.111Z	INFO	Connecting	{"node": "no_liquidity_4", "peer": "destination2", "peerPubKey": "027e4933bc5b4de4f22635c458f4735bb447933a8d6df0af49292e2c69d9bfdead", "host": "node_destination2:9735"}
testrunner_1                | 2022-05-16T08:12:58.119Z	INFO	Open channel	{"node": "no_liquidity_4", "peer": "destination2"}
testrunner_1                | 2022-05-16T08:12:58.347Z	INFO	Connecting	{"node": "start", "peer": "node0_1", "peerPubKey": "036fa23150634644b42442bf375c46e845083cb33ebd9e430dfe8ab877de175296", "host": "node_node0_1:9735"}
testrunner_1                | 2022-05-16T08:12:58.352Z	INFO	Open channel	{"node": "start", "peer": "node0_1"}
testrunner_1                | 2022-05-16T08:12:58.431Z	INFO	Connecting	{"node": "start", "peer": "node0_0", "peerPubKey": "03ce56a2e81b6491420b0063eb11a7e13a7af32fbcb0a348ce4f1d578ad1e25064", "host": "node_node0_0:9735"}
testrunner_1                | 2022-05-16T08:12:58.436Z	INFO	Open channel	{"node": "start", "peer": "node0_0"}
testrunner_1                | 2022-05-16T08:12:58.592Z	INFO	Connecting	{"node": "node2_0", "peer": "node3_0", "peerPubKey": "031ce4fd8535b25aa8c9090c21c282bfce10e3143c5573553c18f4e3fb14d2d1b2", "host": "node_node3_0:9735"}
testrunner_1                | 2022-05-16T08:12:58.597Z	INFO	Open channel	{"node": "node2_0", "peer": "node3_0"}
testrunner_1                | 2022-05-16T08:12:58.737Z	INFO	Connecting	{"node": "node2_0", "peer": "node3_1", "peerPubKey": "027deb0a8ff6a857a51b8c34b1826ef32deb85b2b8843ab83cd7fb8aa8ee84792e", "host": "node_node3_1:9735"}
testrunner_1                | 2022-05-16T08:12:58.743Z	INFO	Open channel	{"node": "node2_0", "peer": "node3_1"}
testrunner_1                | 2022-05-16T08:12:58.888Z	INFO	Connecting	{"node": "node3_0", "peer": "node4_0", "peerPubKey": "0325c9adb14fd2052a0705a8298316fbff71afe3a577d6fcf413c47160596d0b14", "host": "node_node4_0:9735"}
testrunner_1                | 2022-05-16T08:12:58.892Z	INFO	Open channel	{"node": "node3_0", "peer": "node4_0"}
testrunner_1                | 2022-05-16T08:12:58.961Z	INFO	Connecting	{"node": "node3_0", "peer": "node4_1", "peerPubKey": "02da9062200f01617131d75bf45775340323aa31dbb992255fa5513d90f8857787", "host": "node_node4_1:9735"}
testrunner_1                | 2022-05-16T08:12:58.967Z	INFO	Open channel	{"node": "node3_0", "peer": "node4_1"}
testrunner_1                | 2022-05-16T08:12:59.118Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_2", "peerPubKey": "0374e63ee63dcabe31b8875d8d8a0792d2b38918ca5014502f9585a51738e42c58", "host": "node_no_liquidity_2:9735"}
testrunner_1                | 2022-05-16T08:12:59.122Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_2"}
testrunner_1                | 2022-05-16T08:12:59.189Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_3", "peerPubKey": "034c887a1d347eab4e0c98d1a308c2c3a2a0dc46f200e7c9b665d808d9a8d0ace0", "host": "node_no_liquidity_3:9735"}
testrunner_1                | 2022-05-16T08:12:59.195Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_3"}
testrunner_1                | 2022-05-16T08:12:59.258Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_4", "peerPubKey": "024d788fb7f741bf9e8d353819fda5a715284ed6882a5a2485b0463b8c605e0f0e", "host": "node_no_liquidity_4:9735"}
testrunner_1                | 2022-05-16T08:12:59.264Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_4"}
testrunner_1                | 2022-05-16T08:12:59.330Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_5", "peerPubKey": "03db61e38801e90797034d9b2488ec3c6ac7372d85293270f2bdbbdfb085988a6e", "host": "node_no_liquidity_5:9735"}
testrunner_1                | 2022-05-16T08:12:59.334Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_5"}
testrunner_1                | 2022-05-16T08:12:59.401Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_0", "peerPubKey": "030e5e975e438d9c65f663d8bce24f9352466e8e9376727330369da0451f10cfe8", "host": "node_no_liquidity_0:9735"}
testrunner_1                | 2022-05-16T08:12:59.406Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_0"}
testrunner_1                | 2022-05-16T08:12:59.477Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_1", "peerPubKey": "0225087786087a9db4a8b32fd35761babacf7bb8d13732dace57a132ea907c3950", "host": "node_no_liquidity_1:9735"}
testrunner_1                | 2022-05-16T08:12:59.483Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_1"}
testrunner_1                | 2022-05-16T08:12:59.727Z	INFO	Connecting	{"node": "node0_1", "peer": "node1_0", "peerPubKey": "03b0e27c266d3982eca208fc435fb6c0109acc6e9c023dc4664470ffd01da1e3e2", "host": "node_node1_0:9735"}
testrunner_1                | 2022-05-16T08:12:59.731Z	INFO	Open channel	{"node": "node0_1", "peer": "node1_0"}
testrunner_1                | 2022-05-16T08:12:59.801Z	INFO	Connecting	{"node": "node0_1", "peer": "node1_1", "peerPubKey": "03071ce2ce4e66ce7743c80d97a1a98cd94ed5c8a4d8a5aa6c8d691a3a4369c23e", "host": "node_node1_1:9735"}
testrunner_1                | 2022-05-16T08:12:59.807Z	INFO	Open channel	{"node": "node0_1", "peer": "node1_1"}
testrunner_1                | 2022-05-16T08:12:59.957Z	INFO	Connecting	{"node": "node1_0", "peer": "node2_0", "peerPubKey": "02b1a1f35e76116a0460e100f65b952c4ebe6e04c973a610c36f7950a5f91bdd64", "host": "node_node2_0:9735"}
testrunner_1                | 2022-05-16T08:12:59.961Z	INFO	Open channel	{"node": "node1_0", "peer": "node2_0"}
testrunner_1                | 2022-05-16T08:13:00.023Z	INFO	Connecting	{"node": "node1_0", "peer": "node2_1", "peerPubKey": "020d9e75637de128a357844440fa37ff7adb99df4dc89dddd85dfc29dcd2d95c32", "host": "node_node2_1:9735"}
testrunner_1                | 2022-05-16T08:13:00.030Z	INFO	Open channel	{"node": "node1_0", "peer": "node2_1"}
testrunner_1                | 2022-05-16T08:13:00.172Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_2", "peerPubKey": "0374e63ee63dcabe31b8875d8d8a0792d2b38918ca5014502f9585a51738e42c58", "host": "node_no_liquidity_2:9735"}
testrunner_1                | 2022-05-16T08:13:00.177Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_2"}
testrunner_1                | 2022-05-16T08:13:00.246Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_3", "peerPubKey": "034c887a1d347eab4e0c98d1a308c2c3a2a0dc46f200e7c9b665d808d9a8d0ace0", "host": "node_no_liquidity_3:9735"}
testrunner_1                | 2022-05-16T08:13:00.284Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_3"}
testrunner_1                | 2022-05-16T08:13:00.347Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_4", "peerPubKey": "024d788fb7f741bf9e8d353819fda5a715284ed6882a5a2485b0463b8c605e0f0e", "host": "node_no_liquidity_4:9735"}
testrunner_1                | 2022-05-16T08:13:00.354Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_4"}
testrunner_1                | 2022-05-16T08:13:00.418Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_5", "peerPubKey": "03db61e38801e90797034d9b2488ec3c6ac7372d85293270f2bdbbdfb085988a6e", "host": "node_no_liquidity_5:9735"}
testrunner_1                | 2022-05-16T08:13:00.424Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_5"}
testrunner_1                | 2022-05-16T08:13:00.488Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_0", "peerPubKey": "030e5e975e438d9c65f663d8bce24f9352466e8e9376727330369da0451f10cfe8", "host": "node_no_liquidity_0:9735"}
testrunner_1                | 2022-05-16T08:13:00.495Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_0"}
testrunner_1                | 2022-05-16T08:13:00.559Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_1", "peerPubKey": "0225087786087a9db4a8b32fd35761babacf7bb8d13732dace57a132ea907c3950", "host": "node_no_liquidity_1:9735"}
testrunner_1                | 2022-05-16T08:13:00.564Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_1"}
testrunner_1                | 2022-05-16T08:13:00.769Z	INFO	Connecting	{"node": "no_liquidity_0", "peer": "destination1", "peerPubKey": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b", "host": "node_destination1:9735"}
testrunner_1                | 2022-05-16T08:13:00.773Z	INFO	Open channel	{"node": "no_liquidity_0", "peer": "destination1"}
testrunner_1                | 2022-05-16T08:13:00.839Z	INFO	Connecting	{"node": "no_liquidity_0", "peer": "destination2", "peerPubKey": "027e4933bc5b4de4f22635c458f4735bb447933a8d6df0af49292e2c69d9bfdead", "host": "node_destination2:9735"}
testrunner_1                | 2022-05-16T08:13:00.846Z	INFO	Open channel	{"node": "no_liquidity_0", "peer": "destination2"}
testrunner_1                | 2022-05-16T08:13:00.989Z	INFO	Connecting	{"node": "node0_0", "peer": "node1_1", "peerPubKey": "03071ce2ce4e66ce7743c80d97a1a98cd94ed5c8a4d8a5aa6c8d691a3a4369c23e", "host": "node_node1_1:9735"}
testrunner_1                | 2022-05-16T08:13:00.994Z	INFO	Open channel	{"node": "node0_0", "peer": "node1_1"}
testrunner_1                | 2022-05-16T08:13:01.062Z	INFO	Connecting	{"node": "node0_0", "peer": "black_hole", "peerPubKey": "0256bd7a7cd87c5179bcff888704e75ee7e620e042d5d16e861600516c5dd4b71c", "host": "node_black_hole:9735"}
testrunner_1                | 2022-05-16T08:13:01.068Z	INFO	Open channel	{"node": "node0_0", "peer": "black_hole"}
testrunner_1                | 2022-05-16T08:13:01.136Z	INFO	Connecting	{"node": "node0_0", "peer": "black_hole_indirect", "peerPubKey": "02021eb03f5cb8dd028a63e41388acc89f6fece860794aaa3e8ad885acf040c13c", "host": "node_black_hole_indirect:9735"}
testrunner_1                | 2022-05-16T08:13:01.142Z	INFO	Open channel	{"node": "node0_0", "peer": "black_hole_indirect"}
testrunner_1                | 2022-05-16T08:13:01.216Z	INFO	Connecting	{"node": "node0_0", "peer": "node1_0", "peerPubKey": "03b0e27c266d3982eca208fc435fb6c0109acc6e9c023dc4664470ffd01da1e3e2", "host": "node_node1_0:9735"}
testrunner_1                | 2022-05-16T08:13:01.221Z	INFO	Open channel	{"node": "node0_0", "peer": "node1_0"}
testrunner_1                | 2022-05-16T08:13:01.366Z	INFO	Connecting	{"node": "node2_1", "peer": "node3_0", "peerPubKey": "031ce4fd8535b25aa8c9090c21c282bfce10e3143c5573553c18f4e3fb14d2d1b2", "host": "node_node3_0:9735"}
testrunner_1                | 2022-05-16T08:13:01.372Z	INFO	Open channel	{"node": "node2_1", "peer": "node3_0"}
testrunner_1                | 2022-05-16T08:13:01.436Z	INFO	Connecting	{"node": "node2_1", "peer": "node3_1", "peerPubKey": "027deb0a8ff6a857a51b8c34b1826ef32deb85b2b8843ab83cd7fb8aa8ee84792e", "host": "node_node3_1:9735"}
testrunner_1                | 2022-05-16T08:13:01.443Z	INFO	Open channel	{"node": "node2_1", "peer": "node3_1"}
testrunner_1                | 2022-05-16T08:13:01.585Z	INFO	Connecting	{"node": "no_liquidity_1", "peer": "destination1", "peerPubKey": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b", "host": "node_destination1:9735"}
testrunner_1                | 2022-05-16T08:13:01.590Z	INFO	Open channel	{"node": "no_liquidity_1", "peer": "destination1"}
testrunner_1                | 2022-05-16T08:13:01.683Z	INFO	Connecting	{"node": "no_liquidity_1", "peer": "destination2", "peerPubKey": "027e4933bc5b4de4f22635c458f4735bb447933a8d6df0af49292e2c69d9bfdead", "host": "node_destination2:9735"}
testrunner_1                | 2022-05-16T08:13:01.689Z	INFO	Open channel	{"node": "no_liquidity_1", "peer": "destination2"}
testrunner_1                | 2022-05-16T08:13:01.833Z	INFO	Connecting	{"node": "no_liquidity_5", "peer": "destination1", "peerPubKey": "0321fe51da501f6971c4b5e48c42916d82854c4e8dde79161db0ccfa89ba00fb6b", "host": "node_destination1:9735"}
testrunner_1                | 2022-05-16T08:13:01.838Z	INFO	Open channel	{"node": "no_liquidity_5", "peer": "destination1"}
testrunner_1                | 2022-05-16T08:13:01.902Z	INFO	Connecting	{"node": "no_liquidity_5", "peer": "destination2", "peerPubKey": "027e4933bc5b4de4f22635c458f4735bb447933a8d6df0af49292e2c69d9bfdead", "host": "node_destination2:9735"}
testrunner_1                | 2022-05-16T08:13:01.908Z	INFO	Open channel	{"node": "no_liquidity_5", "peer": "destination2"}
testrunner_1                | 2022-05-16T08:13:01.973Z	INFO	Confirm channels
testrunner_1                | 2022-05-16T08:13:02.058Z	DEBUG	Waiting for active channels	{"node": "node2_1", "expected": 4, "count": 0}
testrunner_1                | 2022-05-16T08:13:03.066Z	DEBUG	Waiting for active channels	{"node": "node2_1", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.069Z	DEBUG	Waiting for active channels	{"node": "node4_1", "expected": 3, "count": 3}
testrunner_1                | 2022-05-16T08:13:03.071Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_4", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.095Z	DEBUG	Waiting for active channels	{"node": "node0_0", "expected": 5, "count": 5}
testrunner_1                | 2022-05-16T08:13:03.100Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_5", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.107Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_3", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.113Z	DEBUG	Waiting for active channels	{"node": "node1_1", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.116Z	DEBUG	Waiting for active channels	{"node": "node4_0", "expected": 3, "count": 3}
testrunner_1                | 2022-05-16T08:13:03.119Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_2", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.128Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_1", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.133Z	DEBUG	Waiting for active channels	{"node": "node1_0", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.145Z	DEBUG	Waiting for active channels	{"node": "destination1", "expected": 8, "count": 8}
testrunner_1                | 2022-05-16T08:13:03.148Z	DEBUG	Waiting for active channels	{"node": "node2_0", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.150Z	DEBUG	Waiting for active channels	{"node": "start", "expected": 2, "count": 2}
testrunner_1                | 2022-05-16T08:13:03.153Z	DEBUG	Waiting for active channels	{"node": "node0_1", "expected": 3, "count": 3}
testrunner_1                | 2022-05-16T08:13:03.158Z	DEBUG	Waiting for active channels	{"node": "black_hole_indirect", "expected": 7, "count": 7}
testrunner_1                | 2022-05-16T08:13:03.161Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_0", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.165Z	DEBUG	Waiting for active channels	{"node": "black_hole", "expected": 7, "count": 7}
testrunner_1                | 2022-05-16T08:13:03.175Z	DEBUG	Waiting for active channels	{"node": "destination2", "expected": 7, "count": 7}
testrunner_1                | 2022-05-16T08:13:03.178Z	DEBUG	Waiting for active channels	{"node": "node3_0", "expected": 4, "count": 4}
testrunner_1                | 2022-05-16T08:13:03.181Z	DEBUG	Waiting for active channels	{"node": "node3_1", "expected": 5, "count": 5}
testrunner_1                | 2022-05-16T08:13:03.181Z	INFO	Wait for propagation
testrunner_1                | 2022-05-16T08:13:03.182Z	DEBUG	Gossiped edges	{"count": 0, "expected": 94}
testrunner_1                | 2022-05-16T08:13:04.183Z	DEBUG	Gossiped edges	{"count": 4, "expected": 94}
testrunner_1                | 2022-05-16T08:13:05.185Z	DEBUG	Gossiped edges	{"count": 19, "expected": 94}
testrunner_1                | 2022-05-16T08:13:06.188Z	DEBUG	Gossiped edges	{"count": 38, "expected": 94}
testrunner_1                | 2022-05-16T08:13:07.189Z	DEBUG	Gossiped edges	{"count": 65, "expected": 94}
testrunner_1                | 2022-05-16T08:13:08.191Z	DEBUG	Gossiped edges	{"count": 91, "expected": 94}
testrunner_1                | 2022-05-16T08:13:09.194Z	DEBUG	Gossiped edges	{"count": 94, "expected": 94}
testrunner_1                | 2022-05-16T08:13:09.198Z	INFO	Sending payment	{"invoice": "lnbcrt250u1p3gypu4pp5zxydff5w6alggrf4qp50xag0hn6z8cyw9pm878a82am2r6uhhkqqdqqcqzpgxqyz5vqsp5n6kemgzp6lrwpvxvw792xwu0cfg5hk335zat8wnsu7x0pg7gqngq9qyyssq049mqg4lanra94q9zy7a97jcxdfevvue25hy7cu7dwvl5mnva2cqm96wnz3nn59a4as834sj6myt27vymtjazkea98p6csvj4nsxt7cqee4ls3"}
testrunner_1                | 2022-05-16T08:13:09.228Z	DEBUG	Payment update	{"htlcIdx": 0, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_4", "destination1"]}
testrunner_1                | 2022-05-16T08:13:09.582Z	DEBUG	Payment update	{"htlcIdx": 1, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination1"]}
testrunner_1                | 2022-05-16T08:13:09.948Z	DEBUG	Payment update	{"htlcIdx": 2, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination1"]}
testrunner_1                | 2022-05-16T08:13:10.295Z	DEBUG	Payment update	{"htlcIdx": 3, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination1"]}
testrunner_1                | 2022-05-16T08:13:10.704Z	DEBUG	Payment update	{"htlcIdx": 4, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:11.044Z	DEBUG	Payment update	{"htlcIdx": 5, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination1"]}
testrunner_1                | 2022-05-16T08:13:11.388Z	DEBUG	Payment update	{"htlcIdx": 6, "amt": 25000, "route": ["node0_0", "node1_0", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:11.523Z	DEBUG	Payment update	{"htlcIdx": 7, "amt": 25000, "route": ["node0_0", "node1_1", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:11.810Z	DEBUG	Payment update	{"htlcIdx": 8, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_4", "destination2", "node3_1", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:12.153Z	DEBUG	Payment update	{"htlcIdx": 9, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2", "node3_1", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:12.497Z	DEBUG	Payment update	{"htlcIdx": 10, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2", "node3_1", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:12.869Z	DEBUG	Payment update	{"htlcIdx": 11, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2", "node3_1", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:13.211Z	DEBUG	Payment update	{"htlcIdx": 12, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2", "node3_1", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:13.550Z	DEBUG	Payment update	{"htlcIdx": 13, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2", "node3_1", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:13.945Z	DEBUG	Payment update	{"htlcIdx": 14, "amt": 25000, "route": ["node0_0", "node1_1", "node2_0", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:14.290Z	DEBUG	Payment update	{"htlcIdx": 15, "amt": 25000, "route": ["node0_0", "node1_1", "node0_1", "node1_0", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:14.674Z	DEBUG	Payment update	{"htlcIdx": 16, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination1"]}
testrunner_1                | 2022-05-16T08:13:14.688Z	DEBUG	Payment update	{"htlcIdx": 17, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination1"]}
testrunner_1                | 2022-05-16T08:13:15.046Z	DEBUG	Payment update	{"htlcIdx": 18, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination1"]}
testrunner_1                | 2022-05-16T08:13:15.076Z	DEBUG	Payment update	{"htlcIdx": 19, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_4", "destination1"]}
testrunner_1                | 2022-05-16T08:13:15.366Z	DEBUG	Payment update	{"htlcIdx": 20, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination1"]}
testrunner_1                | 2022-05-16T08:13:15.398Z	DEBUG	Payment update	{"htlcIdx": 21, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:15.723Z	DEBUG	Payment update	{"htlcIdx": 22, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:15.777Z	DEBUG	Payment update	{"htlcIdx": 23, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination1"]}
testrunner_1                | 2022-05-16T08:13:16.047Z	DEBUG	Payment update	{"htlcIdx": 24, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination1"]}
testrunner_1                | 2022-05-16T08:13:16.129Z	DEBUG	Payment update	{"htlcIdx": 25, "amt": 12500, "route": ["node0_0", "node1_1", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:16.383Z	DEBUG	Payment update	{"htlcIdx": 26, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:16.414Z	DEBUG	Payment update	{"htlcIdx": 27, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:16.765Z	DEBUG	Payment update	{"htlcIdx": 28, "amt": 12500, "route": ["node0_0", "node1_0", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:16.804Z	DEBUG	Payment update	{"htlcIdx": 29, "amt": 12500, "route": ["node0_0", "node1_0", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:16.898Z	DEBUG	Payment update	{"htlcIdx": 30, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:16.954Z	DEBUG	Payment update	{"htlcIdx": 31, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:17.331Z	DEBUG	Payment update	{"htlcIdx": 32, "amt": 12500, "route": ["node0_0", "node1_1", "node0_1", "node1_0", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:17.384Z	DEBUG	Payment update	{"htlcIdx": 33, "amt": 12500, "route": ["node0_0", "node1_1", "node0_1", "node1_0", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:17.694Z	DEBUG	Payment update	{"htlcIdx": 34, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node1_0", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:17.748Z	DEBUG	Payment update	{"htlcIdx": 35, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node1_0", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1                | 2022-05-16T08:13:18.136Z	DEBUG	Payment update	{"htlcIdx": 36, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_1", "destination1"]}
testrunner_1                | 2022-05-16T08:13:18.174Z	DEBUG	Payment update	{"htlcIdx": 37, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_1", "destination1"]}
testrunner_1                | 2022-05-16T08:13:18.875Z	INFO	Time elapsed	{"time": "9.677392047s"}
testrunner_1                | 2022-05-16T08:13:18.879Z	INFO	Sending payment	{"invoice": "lnbcrt250u1p3gypu7pp5gqtxmfhjxewjk022vzrj9q4x6asn0vgn40aerk84fn5mwec4uv8sdqqcqzpgxqyz5vqsp5ggu5a9gkszfc6n4c7q0c7ekjcmlle3l8ysux9rge8h8re7f5xuss9qyyssqvgyhzp7n794k43h69cdgmnw57xjs5fwljlmu6w0zekk5g06aust3y8nxfzfv05wxa3zhreynrshnchya40j4t5895fatnteuaaruccsp8e3sxy"}
testrunner_1                | 2022-05-16T08:13:18.907Z	DEBUG	Payment update	{"htlcIdx": 0, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2"]}
testrunner_1                | 2022-05-16T08:13:19.249Z	DEBUG	Payment update	{"htlcIdx": 1, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_4", "destination2"]}
testrunner_1                | 2022-05-16T08:13:19.597Z	DEBUG	Payment update	{"htlcIdx": 2, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2"]}
testrunner_1                | 2022-05-16T08:13:19.944Z	DEBUG	Payment update	{"htlcIdx": 3, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2"]}
testrunner_1                | 2022-05-16T08:13:20.287Z	DEBUG	Payment update	{"htlcIdx": 4, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2"]}
testrunner_1                | 2022-05-16T08:13:20.654Z	DEBUG	Payment update	{"htlcIdx": 5, "amt": 25000, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2"]}
testrunner_1                | 2022-05-16T08:13:21.005Z	DEBUG	Payment update	{"htlcIdx": 6, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2"]}
testrunner_1                | 2022-05-16T08:13:21.018Z	DEBUG	Payment update	{"htlcIdx": 7, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2"]}
testrunner_1                | 2022-05-16T08:13:21.375Z	DEBUG	Payment update	{"htlcIdx": 8, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_4", "destination2"]}
testrunner_1                | 2022-05-16T08:13:21.400Z	DEBUG	Payment update	{"htlcIdx": 9, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_4", "destination2"]}
testrunner_1                | 2022-05-16T08:13:21.799Z	DEBUG	Payment update	{"htlcIdx": 10, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2"]}
testrunner_1                | 2022-05-16T08:13:21.834Z	DEBUG	Payment update	{"htlcIdx": 11, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2"]}
testrunner_1                | 2022-05-16T08:13:22.127Z	DEBUG	Payment update	{"htlcIdx": 12, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2"]}
testrunner_1                | 2022-05-16T08:13:22.160Z	DEBUG	Payment update	{"htlcIdx": 13, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2"]}
testrunner_1                | 2022-05-16T08:13:22.443Z	DEBUG	Payment update	{"htlcIdx": 14, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2"]}
testrunner_1                | 2022-05-16T08:13:22.477Z	DEBUG	Payment update	{"htlcIdx": 15, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2"]}
testrunner_1                | 2022-05-16T08:13:22.784Z	DEBUG	Payment update	{"htlcIdx": 16, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2"]}
testrunner_1                | 2022-05-16T08:13:22.819Z	DEBUG	Payment update	{"htlcIdx": 17, "amt": 12500, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2"]}
testrunner_1                | 2022-05-16T08:13:23.157Z	DEBUG	Payment update	{"htlcIdx": 18, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "destination2"]}
testrunner_1                | 2022-05-16T08:13:23.188Z	DEBUG	Payment update	{"htlcIdx": 19, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "destination2"]}
testrunner_1                | 2022-05-16T08:13:23.760Z	INFO	Time elapsed	{"time": "4.881456878s"}
testrunner_1                | 2022-05-16T08:13:23.763Z	INFO	Sending payment	{"invoice": "lnbcrt250u1p3gyparpp5agkehlj8253vufd6jutl4etaegnsdrdwm74jy3k5tgg5ye6v54ysdqqcqzpgxqyz5vqsp5cgta8xdswrre8ks47r4rf2fzhfntr22eal7j5r4w6xg7vcel3ras9qyyssqpt9ulrc82krr2vtt9fdjsfer8q54ka2zksqawl5360x3waec4658064jcay0mct4ehrke97wwjkdnazdxgg2txrnacpsflsnpxc9tkcphvmj3d"}
testrunner_1                | 2022-05-16T08:13:23.793Z	DEBUG	Payment update	{"htlcIdx": 0, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_1", "destination1"]}
testrunner_1                | 2022-05-16T08:13:23.810Z	DEBUG	Payment update	{"htlcIdx": 1, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_1", "destination1"]}
testrunner_1                | 2022-05-16T08:13:24.426Z	INFO	Time elapsed	{"time": "662.850708ms"}
testrunner_1                | 2022-05-16T08:13:24.431Z	INFO	Sending payment	{"invoice": "lnbcrt250u1p3gypaypp5svxg7yvyr0c5mlr5a6lun2mt975chtmem56955ef3pdlaclrtwaqdqqcqzpgxqyz5vqsp5347pw6ryxp499kk0c8qg4p4p87xqc88mg8pa85ezd78snl8gvxlq9qyyssqlvwgvav3p6m79q67k8sm3pjpfe3fdemlqm4s78w5fey7f7h43zapa86pyv0zn2uyajz4h64g43gsa6uw2x9y8t4233c49lzjds0gkyspfcrymn"}
testrunner_1                | 2022-05-16T08:13:24.460Z	DEBUG	Payment update	{"htlcIdx": 0, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "destination2"]}
testrunner_1                | 2022-05-16T08:13:24.476Z	DEBUG	Payment update	{"htlcIdx": 1, "amt": 12500, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "destination2"]}
testrunner_1                | 2022-05-16T08:13:25.044Z	INFO	Time elapsed	{"time": "612.840376ms"}
testrunner_1                | 2022-05-16T08:13:25.045Z	INFO	Total time elapsed	{"time": "15.850126799s"}
testrunner_1                | 2022-05-16T08:13:25.045Z	INFO	Done
```