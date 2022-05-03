# Pathfinding Benchmark

The goal of the code in this repository is to provide a node-implementation
independent benchmark for pathfinding algorithms.

Relevant files:

* [graph.yml](graph.yml) defines the test graph. The node `start` is special
  because this is where payments are made from and for which the implementation
  can be switched between CLN and LND. All other nodes in the test graph are LND
  based.

* [cmd/gencluster](cmd/gencluster) contains code to generate the
  [docker-compose](docker-compose.yml) file from `graph.yml`.

* [cmd/testrunner/cmd_run.go](cmd/testrunner/cmd_run.go) is the automated test
  that sets up channels and executes the payments.

* [run.sh](run.sh) fires it all up. The `TARGET` env var controls the
  implementation that is used for the `start` node.

  If all goes well, output should be similar to:

```
testrunner_1               | Starting test.
testrunner_1               | 2022-05-03T20:26:06.631Z	INFO	Attempting to connect to bitcoind
testrunner_1               | 2022-05-03T20:26:06.682Z	INFO	Connected to bitcoind
testrunner_1               | 2022-05-03T20:26:06.683Z	INFO	Creating bitcoind wallet
testrunner_1               | 2022-05-03T20:26:15.817Z	INFO	Bitcoin address	{"address": "bcrt1qlvasm7hxgy80lhl4hh2n75yp8nrkknqvx5zg3k"}
testrunner_1               | 2022-05-03T20:26:15.821Z	INFO	Activate segwit
testrunner_1               | 2022-05-03T20:26:26.968Z	INFO	Fund senders
testrunner_1               | 2022-05-03T20:26:26.971Z	INFO	Attempting to connect to lnd	{"node": "node1_1"}
testrunner_1               | 2022-05-03T20:26:33.377Z	INFO	Connected to lnd	{"node": "node1_1", "key": "033d4794a4e8e3209673511451c82acd293a920e47ff690c6c163290948f80ff97"}
testrunner_1               | 2022-05-03T20:26:33.382Z	INFO	Generated funding address	{"node": "node1_1", "address": "bcrt1qyrmkthwvczux3k4ku3gp4a4z2c0fjxsq5425pn"}
testrunner_1               | 2022-05-03T20:26:34.999Z	INFO	Attempting to connect to lnd	{"node": "node2_1"}
testrunner_1               | 2022-05-03T20:26:35.040Z	INFO	Connected to lnd	{"node": "node2_1", "key": "03d3db72af290ad7d4c20b0786c24b69df7a6829580baf3dc61f3fbb3fac66e8bb"}
testrunner_1               | 2022-05-03T20:26:35.046Z	INFO	Generated funding address	{"node": "node2_1", "address": "bcrt1qh5uvc20760ztlma57l7vhfe6vvaun3xxm43tuy"}
testrunner_1               | 2022-05-03T20:26:36.460Z	INFO	Attempting to connect to lnd	{"node": "node3_1"}
testrunner_1               | 2022-05-03T20:26:36.473Z	INFO	Connected to lnd	{"node": "node3_1", "key": "030dce4716befd19c3f73deb896ba89af38dcbbd00e02a16bb11bdc64ea20a4727"}
testrunner_1               | 2022-05-03T20:26:36.477Z	INFO	Generated funding address	{"node": "node3_1", "address": "bcrt1qvtyqv6ftcjssje4le99040kpt7yyzt4vle0my5"}
testrunner_1               | 2022-05-03T20:26:37.503Z	INFO	Attempting to connect to lnd	{"node": "black_hole"}
testrunner_1               | 2022-05-03T20:26:37.517Z	INFO	Connected to lnd	{"node": "black_hole", "key": "030c90b067162358d4e665647674e26c3c716a0fbcb67d2ef3d7f960771db1dc1e"}
testrunner_1               | 2022-05-03T20:26:37.520Z	INFO	Generated funding address	{"node": "black_hole", "address": "bcrt1q4ppgy8vqzsskwzfuw047u0zehhlcmhn69e57am"}
testrunner_1               | 2022-05-03T20:26:38.553Z	INFO	Attempting to connect to lnd	{"node": "destination1"}
testrunner_1               | 2022-05-03T20:26:38.569Z	INFO	Connected to lnd	{"node": "destination1", "key": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48"}
testrunner_1               | 2022-05-03T20:26:38.617Z	INFO	Generated funding address	{"node": "destination1", "address": "bcrt1qkc4rgpyh9mqpvu0cfl0mzuu4cu7h89vjyxpnf3"}
testrunner_1               | 2022-05-03T20:26:39.634Z	INFO	Attempting to connect to lnd	{"node": "start"}
testrunner_1               | 2022-05-03T20:26:39.650Z	INFO	Connected to lnd	{"node": "start", "key": "03938976873dc999f0a931de9c8e8b2adbde7bddf7d36f01d94f27561012e35b36"}
testrunner_1               | 2022-05-03T20:26:39.706Z	INFO	Generated funding address	{"node": "start", "address": "bcrt1qazdmcphr8wrakzfz5kxdwyh6q9l44rj6ep3jf2"}
testrunner_1               | 2022-05-03T20:26:40.710Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_0"}
testrunner_1               | 2022-05-03T20:26:40.724Z	INFO	Connected to lnd	{"node": "no_liquidity_0", "key": "0393dec60b23e9b9a2d3d83897faab922451e2b54877ff617a7bdd981ddb213c17"}
testrunner_1               | 2022-05-03T20:26:40.726Z	INFO	Generated funding address	{"node": "no_liquidity_0", "address": "bcrt1qygu83fhc9caj2a22mlmvhkjtmrjfsfnlqxesse"}
testrunner_1               | 2022-05-03T20:26:41.738Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_2"}
testrunner_1               | 2022-05-03T20:26:41.752Z	INFO	Connected to lnd	{"node": "no_liquidity_2", "key": "03ce4f58635c5dc96cb7b4407f4144a77f92300a0096b82dfb1107b91079b39fd2"}
testrunner_1               | 2022-05-03T20:26:41.756Z	INFO	Generated funding address	{"node": "no_liquidity_2", "address": "bcrt1qsday42xtzs3e9wpljep92mnqrtta0gm7u6r0v2"}
testrunner_1               | 2022-05-03T20:26:42.755Z	INFO	Attempting to connect to lnd	{"node": "black_hole_indirect"}
testrunner_1               | 2022-05-03T20:26:42.770Z	INFO	Connected to lnd	{"node": "black_hole_indirect", "key": "02fd6cdbe61b2b5268a846216dea94d3665c79d32d6b851cd05624dd50039972b5"}
testrunner_1               | 2022-05-03T20:26:42.774Z	INFO	Generated funding address	{"node": "black_hole_indirect", "address": "bcrt1qtkmxknqg5px5ke45seypuuymjlz866frhnlxex"}
testrunner_1               | 2022-05-03T20:26:43.789Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_4"}
testrunner_1               | 2022-05-03T20:26:43.804Z	INFO	Connected to lnd	{"node": "no_liquidity_4", "key": "0342ef38c0a7950539301418ab786ced22a5aaeb3d91cd620322b51b7025cffcf4"}
testrunner_1               | 2022-05-03T20:26:43.808Z	INFO	Generated funding address	{"node": "no_liquidity_4", "address": "bcrt1q2sxhwt4fjyd3fz8tvxmcj52jnty32dfytsj5z5"}
testrunner_1               | 2022-05-03T20:26:44.814Z	INFO	Attempting to connect to lnd	{"node": "destination2"}
testrunner_1               | 2022-05-03T20:26:44.826Z	INFO	Connected to lnd	{"node": "destination2", "key": "03934d97ef56d7f5b7ec6dfabe25fac0852e10a2d86fe7ec1846efa6784e568e30"}
testrunner_1               | 2022-05-03T20:26:44.830Z	INFO	Generated funding address	{"node": "destination2", "address": "bcrt1qzunqngseazn5h90smatm7y8775slqtjrzuzxw0"}
testrunner_1               | 2022-05-03T20:26:45.858Z	INFO	Attempting to connect to lnd	{"node": "node0_0"}
testrunner_1               | 2022-05-03T20:26:45.874Z	INFO	Connected to lnd	{"node": "node0_0", "key": "02eaa656cc1f1e89c5866a1614553cd61a88c24923b2756febc855052d3f8ba5d2"}
testrunner_1               | 2022-05-03T20:26:45.878Z	INFO	Generated funding address	{"node": "node0_0", "address": "bcrt1qwmnm4vthk8kjlfnv2ykkvakn53mwhf7rnfpzle"}
testrunner_1               | 2022-05-03T20:26:46.895Z	INFO	Attempting to connect to lnd	{"node": "node1_0"}
testrunner_1               | 2022-05-03T20:26:46.909Z	INFO	Connected to lnd	{"node": "node1_0", "key": "02ee621d38fcfc6e89cfc64e3673f6bb8941686062a352deec893bcc48c0e3c56f"}
testrunner_1               | 2022-05-03T20:26:46.912Z	INFO	Generated funding address	{"node": "node1_0", "address": "bcrt1qyjyy2yavzmsc8agq3mcwcunqg3gn0x9wwrpw4l"}
testrunner_1               | 2022-05-03T20:26:47.929Z	INFO	Attempting to connect to lnd	{"node": "node2_0"}
testrunner_1               | 2022-05-03T20:26:47.946Z	INFO	Connected to lnd	{"node": "node2_0", "key": "02e86b182e66c6c50a0fba932d4b6db62f100bb3ad4c1f1e6ecb6d49369265e529"}
testrunner_1               | 2022-05-03T20:26:47.949Z	INFO	Generated funding address	{"node": "node2_0", "address": "bcrt1qatx4vjw4rz2anaduvyjjceufcnrjdvxcpkv6qt"}
testrunner_1               | 2022-05-03T20:26:48.995Z	INFO	Attempting to connect to lnd	{"node": "node4_1"}
testrunner_1               | 2022-05-03T20:26:49.008Z	INFO	Connected to lnd	{"node": "node4_1", "key": "036941506a6c01acdccfff0bfc0c1e45163339427a75b73ee970de48a5c2193831"}
testrunner_1               | 2022-05-03T20:26:49.012Z	INFO	Generated funding address	{"node": "node4_1", "address": "bcrt1qvtju0hextt959h42jqmngns9vewa959kxferzd"}
testrunner_1               | 2022-05-03T20:26:50.058Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_5"}
testrunner_1               | 2022-05-03T20:26:50.072Z	INFO	Connected to lnd	{"node": "no_liquidity_5", "key": "037dff80421d779975f62a24cef168f76c1506611913a88589e13956bd6a7dff4f"}
testrunner_1               | 2022-05-03T20:26:50.077Z	INFO	Generated funding address	{"node": "no_liquidity_5", "address": "bcrt1qyfa5u5nswguxxvt7073xlwhq8585dnvf4esa6k"}
testrunner_1               | 2022-05-03T20:26:51.183Z	INFO	Attempting to connect to lnd	{"node": "node0_1"}
testrunner_1               | 2022-05-03T20:26:51.195Z	INFO	Connected to lnd	{"node": "node0_1", "key": "03349d363fddf9d019f020f0a8785c5874ead5310a1f569ceb9bc4cba21a299cd6"}
testrunner_1               | 2022-05-03T20:26:51.199Z	INFO	Generated funding address	{"node": "node0_1", "address": "bcrt1qv3702n5p85jyetpyq23vpf6gckghh5acla47hr"}
testrunner_1               | 2022-05-03T20:26:52.213Z	INFO	Attempting to connect to lnd	{"node": "node3_0"}
testrunner_1               | 2022-05-03T20:26:52.228Z	INFO	Connected to lnd	{"node": "node3_0", "key": "02d0370de1f5c969ed6be0a734e3964aa30c989a358885193ace0d824621f00e3b"}
testrunner_1               | 2022-05-03T20:26:52.232Z	INFO	Generated funding address	{"node": "node3_0", "address": "bcrt1qrvnd9n2l0tth2qqfqx89xrx0g48s6737fkjgsj"}
testrunner_1               | 2022-05-03T20:26:53.245Z	INFO	Attempting to connect to lnd	{"node": "node4_0"}
testrunner_1               | 2022-05-03T20:26:53.259Z	INFO	Connected to lnd	{"node": "node4_0", "key": "02f2e44ce1b0e8ee8316de8982803c9f0513021e7430e3b1beb04e2dc799551284"}
testrunner_1               | 2022-05-03T20:26:53.263Z	INFO	Generated funding address	{"node": "node4_0", "address": "bcrt1qwwet56ghg83agjal45h24cfkcqum0vxv5cq9k4"}
testrunner_1               | 2022-05-03T20:26:54.277Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_1"}
testrunner_1               | 2022-05-03T20:26:54.291Z	INFO	Connected to lnd	{"node": "no_liquidity_1", "key": "03b92a67d25cacf8e76c5472610534842686e729812d975728b82d7a922a1652c9"}
testrunner_1               | 2022-05-03T20:26:54.295Z	INFO	Generated funding address	{"node": "no_liquidity_1", "address": "bcrt1q2gp6fg3xq2q64dwmwlvf3zs7j3f6rptalufqm5"}
testrunner_1               | 2022-05-03T20:26:55.296Z	INFO	Attempting to connect to lnd	{"node": "no_liquidity_3"}
testrunner_1               | 2022-05-03T20:26:55.319Z	INFO	Connected to lnd	{"node": "no_liquidity_3", "key": "037a72142f30d57da5c6dbe5bbf45c4c7df5c4516519695b5c54a5caaa9e2ce73f"}
testrunner_1               | 2022-05-03T20:26:55.324Z	INFO	Generated funding address	{"node": "no_liquidity_3", "address": "bcrt1qc097x6cwqhmg3jvzpj3gpmywu5dnkny3xs535m"}
testrunner_1               | 2022-05-03T20:26:56.361Z	INFO	Wait for coin and open channels
testrunner_1               | 2022-05-03T20:26:57.459Z	INFO	Connecting	{"node": "node4_0", "peer": "destination1", "peerPubKey": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48", "host": "lnd_destination1:9735"}
testrunner_1               | 2022-05-03T20:26:57.463Z	INFO	Open channel	{"node": "node4_0", "peer": "destination1"}
testrunner_1               | 2022-05-03T20:26:57.627Z	INFO	Connecting	{"node": "no_liquidity_1", "peer": "destination1", "peerPubKey": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48", "host": "lnd_destination1:9735"}
testrunner_1               | 2022-05-03T20:26:57.632Z	INFO	Open channel	{"node": "no_liquidity_1", "peer": "destination1"}
testrunner_1               | 2022-05-03T20:26:57.752Z	INFO	Connecting	{"node": "no_liquidity_1", "peer": "destination2", "peerPubKey": "03934d97ef56d7f5b7ec6dfabe25fac0852e10a2d86fe7ec1846efa6784e568e30", "host": "lnd_destination2:9735"}
testrunner_1               | 2022-05-03T20:26:57.758Z	INFO	Open channel	{"node": "no_liquidity_1", "peer": "destination2"}
testrunner_1               | 2022-05-03T20:26:57.949Z	INFO	Connecting	{"node": "no_liquidity_3", "peer": "destination1", "peerPubKey": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48", "host": "lnd_destination1:9735"}
testrunner_1               | 2022-05-03T20:26:57.955Z	INFO	Open channel	{"node": "no_liquidity_3", "peer": "destination1"}
testrunner_1               | 2022-05-03T20:26:58.025Z	INFO	Connecting	{"node": "no_liquidity_3", "peer": "destination2", "peerPubKey": "03934d97ef56d7f5b7ec6dfabe25fac0852e10a2d86fe7ec1846efa6784e568e30", "host": "lnd_destination2:9735"}
testrunner_1               | 2022-05-03T20:26:58.031Z	INFO	Open channel	{"node": "no_liquidity_3", "peer": "destination2"}
testrunner_1               | 2022-05-03T20:26:58.178Z	INFO	Connecting	{"node": "no_liquidity_5", "peer": "destination1", "peerPubKey": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48", "host": "lnd_destination1:9735"}
testrunner_1               | 2022-05-03T20:26:58.183Z	INFO	Open channel	{"node": "no_liquidity_5", "peer": "destination1"}
testrunner_1               | 2022-05-03T20:26:58.255Z	INFO	Connecting	{"node": "no_liquidity_5", "peer": "destination2", "peerPubKey": "03934d97ef56d7f5b7ec6dfabe25fac0852e10a2d86fe7ec1846efa6784e568e30", "host": "lnd_destination2:9735"}
testrunner_1               | 2022-05-03T20:26:58.260Z	INFO	Open channel	{"node": "no_liquidity_5", "peer": "destination2"}
testrunner_1               | 2022-05-03T20:26:58.410Z	INFO	Connecting	{"node": "node0_1", "peer": "node1_0", "peerPubKey": "02ee621d38fcfc6e89cfc64e3673f6bb8941686062a352deec893bcc48c0e3c56f", "host": "lnd_node1_0:9735"}
testrunner_1               | 2022-05-03T20:26:58.415Z	INFO	Open channel	{"node": "node0_1", "peer": "node1_0"}
testrunner_1               | 2022-05-03T20:26:58.485Z	INFO	Connecting	{"node": "node0_1", "peer": "node1_1", "peerPubKey": "033d4794a4e8e3209673511451c82acd293a920e47ff690c6c163290948f80ff97", "host": "lnd_node1_1:9735"}
testrunner_1               | 2022-05-03T20:26:58.492Z	INFO	Open channel	{"node": "node0_1", "peer": "node1_1"}
testrunner_1               | 2022-05-03T20:26:58.660Z	INFO	Connecting	{"node": "node3_0", "peer": "node4_0", "peerPubKey": "02f2e44ce1b0e8ee8316de8982803c9f0513021e7430e3b1beb04e2dc799551284", "host": "lnd_node4_0:9735"}
testrunner_1               | 2022-05-03T20:26:58.673Z	INFO	Open channel	{"node": "node3_0", "peer": "node4_0"}
testrunner_1               | 2022-05-03T20:26:58.747Z	INFO	Connecting	{"node": "node3_0", "peer": "node4_1", "peerPubKey": "036941506a6c01acdccfff0bfc0c1e45163339427a75b73ee970de48a5c2193831", "host": "lnd_node4_1:9735"}
testrunner_1               | 2022-05-03T20:26:58.753Z	INFO	Open channel	{"node": "node3_0", "peer": "node4_1"}
testrunner_1               | 2022-05-03T20:26:58.903Z	INFO	Connecting	{"node": "node3_1", "peer": "node4_1", "peerPubKey": "036941506a6c01acdccfff0bfc0c1e45163339427a75b73ee970de48a5c2193831", "host": "lnd_node4_1:9735"}
testrunner_1               | 2022-05-03T20:26:58.907Z	INFO	Open channel	{"node": "node3_1", "peer": "node4_1"}
testrunner_1               | 2022-05-03T20:26:58.974Z	INFO	Connecting	{"node": "node3_1", "peer": "destination2", "peerPubKey": "03934d97ef56d7f5b7ec6dfabe25fac0852e10a2d86fe7ec1846efa6784e568e30", "host": "lnd_destination2:9735"}
testrunner_1               | 2022-05-03T20:26:58.980Z	INFO	Open channel	{"node": "node3_1", "peer": "destination2"}
testrunner_1               | 2022-05-03T20:26:59.050Z	INFO	Connecting	{"node": "node3_1", "peer": "node4_0", "peerPubKey": "02f2e44ce1b0e8ee8316de8982803c9f0513021e7430e3b1beb04e2dc799551284", "host": "lnd_node4_0:9735"}
testrunner_1               | 2022-05-03T20:26:59.056Z	INFO	Open channel	{"node": "node3_1", "peer": "node4_0"}
testrunner_1               | 2022-05-03T20:26:59.214Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_0", "peerPubKey": "0393dec60b23e9b9a2d3d83897faab922451e2b54877ff617a7bdd981ddb213c17", "host": "lnd_no_liquidity_0:9735"}
testrunner_1               | 2022-05-03T20:26:59.221Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_0"}
testrunner_1               | 2022-05-03T20:26:59.314Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_1", "peerPubKey": "03b92a67d25cacf8e76c5472610534842686e729812d975728b82d7a922a1652c9", "host": "lnd_no_liquidity_1:9735"}
testrunner_1               | 2022-05-03T20:26:59.320Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_1"}
testrunner_1               | 2022-05-03T20:26:59.388Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_2", "peerPubKey": "03ce4f58635c5dc96cb7b4407f4144a77f92300a0096b82dfb1107b91079b39fd2", "host": "lnd_no_liquidity_2:9735"}
testrunner_1               | 2022-05-03T20:26:59.394Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_2"}
testrunner_1               | 2022-05-03T20:26:59.466Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_3", "peerPubKey": "037a72142f30d57da5c6dbe5bbf45c4c7df5c4516519695b5c54a5caaa9e2ce73f", "host": "lnd_no_liquidity_3:9735"}
testrunner_1               | 2022-05-03T20:26:59.473Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_3"}
testrunner_1               | 2022-05-03T20:26:59.544Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_4", "peerPubKey": "0342ef38c0a7950539301418ab786ced22a5aaeb3d91cd620322b51b7025cffcf4", "host": "lnd_no_liquidity_4:9735"}
testrunner_1               | 2022-05-03T20:26:59.550Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_4"}
testrunner_1               | 2022-05-03T20:26:59.753Z	INFO	Connecting	{"node": "black_hole", "peer": "no_liquidity_5", "peerPubKey": "037dff80421d779975f62a24cef168f76c1506611913a88589e13956bd6a7dff4f", "host": "lnd_no_liquidity_5:9735"}
testrunner_1               | 2022-05-03T20:26:59.758Z	INFO	Open channel	{"node": "black_hole", "peer": "no_liquidity_5"}
testrunner_1               | 2022-05-03T20:26:59.909Z	INFO	Connecting	{"node": "node1_1", "peer": "node2_0", "peerPubKey": "02e86b182e66c6c50a0fba932d4b6db62f100bb3ad4c1f1e6ecb6d49369265e529", "host": "lnd_node2_0:9735"}
testrunner_1               | 2022-05-03T20:26:59.913Z	INFO	Open channel	{"node": "node1_1", "peer": "node2_0"}
testrunner_1               | 2022-05-03T20:26:59.982Z	INFO	Connecting	{"node": "node1_1", "peer": "node2_1", "peerPubKey": "03d3db72af290ad7d4c20b0786c24b69df7a6829580baf3dc61f3fbb3fac66e8bb", "host": "lnd_node2_1:9735"}
testrunner_1               | 2022-05-03T20:26:59.987Z	INFO	Open channel	{"node": "node1_1", "peer": "node2_1"}
testrunner_1               | 2022-05-03T20:27:00.147Z	INFO	Connecting	{"node": "node2_1", "peer": "node3_1", "peerPubKey": "030dce4716befd19c3f73deb896ba89af38dcbbd00e02a16bb11bdc64ea20a4727", "host": "lnd_node3_1:9735"}
testrunner_1               | 2022-05-03T20:27:00.152Z	INFO	Open channel	{"node": "node2_1", "peer": "node3_1"}
testrunner_1               | 2022-05-03T20:27:00.218Z	INFO	Connecting	{"node": "node2_1", "peer": "node3_0", "peerPubKey": "02d0370de1f5c969ed6be0a734e3964aa30c989a358885193ace0d824621f00e3b", "host": "lnd_node3_0:9735"}
testrunner_1               | 2022-05-03T20:27:00.224Z	INFO	Open channel	{"node": "node2_1", "peer": "node3_0"}
testrunner_1               | 2022-05-03T20:27:00.381Z	INFO	Connecting	{"node": "no_liquidity_0", "peer": "destination1", "peerPubKey": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48", "host": "lnd_destination1:9735"}
testrunner_1               | 2022-05-03T20:27:00.386Z	INFO	Open channel	{"node": "no_liquidity_0", "peer": "destination1"}
testrunner_1               | 2022-05-03T20:27:00.451Z	INFO	Connecting	{"node": "no_liquidity_0", "peer": "destination2", "peerPubKey": "03934d97ef56d7f5b7ec6dfabe25fac0852e10a2d86fe7ec1846efa6784e568e30", "host": "lnd_destination2:9735"}
testrunner_1               | 2022-05-03T20:27:00.458Z	INFO	Open channel	{"node": "no_liquidity_0", "peer": "destination2"}
testrunner_1               | 2022-05-03T20:27:00.606Z	INFO	Connecting	{"node": "no_liquidity_2", "peer": "destination1", "peerPubKey": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48", "host": "lnd_destination1:9735"}
testrunner_1               | 2022-05-03T20:27:00.635Z	INFO	Open channel	{"node": "no_liquidity_2", "peer": "destination1"}
testrunner_1               | 2022-05-03T20:27:00.748Z	INFO	Connecting	{"node": "no_liquidity_2", "peer": "destination2", "peerPubKey": "03934d97ef56d7f5b7ec6dfabe25fac0852e10a2d86fe7ec1846efa6784e568e30", "host": "lnd_destination2:9735"}
testrunner_1               | 2022-05-03T20:27:00.755Z	INFO	Open channel	{"node": "no_liquidity_2", "peer": "destination2"}
testrunner_1               | 2022-05-03T20:27:00.987Z	INFO	Connecting	{"node": "start", "peer": "node0_0", "peerPubKey": "02eaa656cc1f1e89c5866a1614553cd61a88c24923b2756febc855052d3f8ba5d2", "host": "lnd_node0_0:9735"}
testrunner_1               | 2022-05-03T20:27:00.994Z	INFO	Open channel	{"node": "start", "peer": "node0_0"}
testrunner_1               | 2022-05-03T20:27:01.066Z	INFO	Connecting	{"node": "start", "peer": "node0_1", "peerPubKey": "03349d363fddf9d019f020f0a8785c5874ead5310a1f569ceb9bc4cba21a299cd6", "host": "lnd_node0_1:9735"}
testrunner_1               | 2022-05-03T20:27:01.074Z	INFO	Open channel	{"node": "start", "peer": "node0_1"}
testrunner_1               | 2022-05-03T20:27:01.220Z	INFO	Connecting	{"node": "node1_0", "peer": "node2_0", "peerPubKey": "02e86b182e66c6c50a0fba932d4b6db62f100bb3ad4c1f1e6ecb6d49369265e529", "host": "lnd_node2_0:9735"}
testrunner_1               | 2022-05-03T20:27:01.224Z	INFO	Open channel	{"node": "node1_0", "peer": "node2_0"}
testrunner_1               | 2022-05-03T20:27:01.292Z	INFO	Connecting	{"node": "node1_0", "peer": "node2_1", "peerPubKey": "03d3db72af290ad7d4c20b0786c24b69df7a6829580baf3dc61f3fbb3fac66e8bb", "host": "lnd_node2_1:9735"}
testrunner_1               | 2022-05-03T20:27:01.299Z	INFO	Open channel	{"node": "node1_0", "peer": "node2_1"}
testrunner_1               | 2022-05-03T20:27:01.446Z	INFO	Connecting	{"node": "node2_0", "peer": "node3_0", "peerPubKey": "02d0370de1f5c969ed6be0a734e3964aa30c989a358885193ace0d824621f00e3b", "host": "lnd_node3_0:9735"}
testrunner_1               | 2022-05-03T20:27:01.451Z	INFO	Open channel	{"node": "node2_0", "peer": "node3_0"}
testrunner_1               | 2022-05-03T20:27:01.519Z	INFO	Connecting	{"node": "node2_0", "peer": "node3_1", "peerPubKey": "030dce4716befd19c3f73deb896ba89af38dcbbd00e02a16bb11bdc64ea20a4727", "host": "lnd_node3_1:9735"}
testrunner_1               | 2022-05-03T20:27:01.525Z	INFO	Open channel	{"node": "node2_0", "peer": "node3_1"}
testrunner_1               | 2022-05-03T20:27:01.777Z	INFO	Connecting	{"node": "node4_1", "peer": "destination1", "peerPubKey": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48", "host": "lnd_destination1:9735"}
testrunner_1               | 2022-05-03T20:27:01.782Z	INFO	Open channel	{"node": "node4_1", "peer": "destination1"}
testrunner_1               | 2022-05-03T20:27:01.934Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_3", "peerPubKey": "037a72142f30d57da5c6dbe5bbf45c4c7df5c4516519695b5c54a5caaa9e2ce73f", "host": "lnd_no_liquidity_3:9735"}
testrunner_1               | 2022-05-03T20:27:01.940Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_3"}
testrunner_1               | 2022-05-03T20:27:02.008Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_4", "peerPubKey": "0342ef38c0a7950539301418ab786ced22a5aaeb3d91cd620322b51b7025cffcf4", "host": "lnd_no_liquidity_4:9735"}
testrunner_1               | 2022-05-03T20:27:02.018Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_4"}
testrunner_1               | 2022-05-03T20:27:02.083Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_5", "peerPubKey": "037dff80421d779975f62a24cef168f76c1506611913a88589e13956bd6a7dff4f", "host": "lnd_no_liquidity_5:9735"}
testrunner_1               | 2022-05-03T20:27:02.088Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_5"}
testrunner_1               | 2022-05-03T20:27:02.154Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_0", "peerPubKey": "0393dec60b23e9b9a2d3d83897faab922451e2b54877ff617a7bdd981ddb213c17", "host": "lnd_no_liquidity_0:9735"}
testrunner_1               | 2022-05-03T20:27:02.161Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_0"}
testrunner_1               | 2022-05-03T20:27:02.228Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_1", "peerPubKey": "03b92a67d25cacf8e76c5472610534842686e729812d975728b82d7a922a1652c9", "host": "lnd_no_liquidity_1:9735"}
testrunner_1               | 2022-05-03T20:27:02.235Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_1"}
testrunner_1               | 2022-05-03T20:27:02.307Z	INFO	Connecting	{"node": "black_hole_indirect", "peer": "no_liquidity_2", "peerPubKey": "03ce4f58635c5dc96cb7b4407f4144a77f92300a0096b82dfb1107b91079b39fd2", "host": "lnd_no_liquidity_2:9735"}
testrunner_1               | 2022-05-03T20:27:02.312Z	INFO	Open channel	{"node": "black_hole_indirect", "peer": "no_liquidity_2"}
testrunner_1               | 2022-05-03T20:27:02.464Z	INFO	Connecting	{"node": "no_liquidity_4", "peer": "destination1", "peerPubKey": "026d60d9efe1ba40874044e56e1c5711e2d6f5e4c35816d9a46c9a85e9a213be48", "host": "lnd_destination1:9735"}
testrunner_1               | 2022-05-03T20:27:02.469Z	INFO	Open channel	{"node": "no_liquidity_4", "peer": "destination1"}
testrunner_1               | 2022-05-03T20:27:02.541Z	INFO	Connecting	{"node": "no_liquidity_4", "peer": "destination2", "peerPubKey": "03934d97ef56d7f5b7ec6dfabe25fac0852e10a2d86fe7ec1846efa6784e568e30", "host": "lnd_destination2:9735"}
testrunner_1               | 2022-05-03T20:27:02.546Z	INFO	Open channel	{"node": "no_liquidity_4", "peer": "destination2"}
testrunner_1               | 2022-05-03T20:27:02.878Z	INFO	Connecting	{"node": "node0_0", "peer": "node1_0", "peerPubKey": "02ee621d38fcfc6e89cfc64e3673f6bb8941686062a352deec893bcc48c0e3c56f", "host": "lnd_node1_0:9735"}
testrunner_1               | 2022-05-03T20:27:02.884Z	INFO	Open channel	{"node": "node0_0", "peer": "node1_0"}
testrunner_1               | 2022-05-03T20:27:02.951Z	INFO	Connecting	{"node": "node0_0", "peer": "node1_1", "peerPubKey": "033d4794a4e8e3209673511451c82acd293a920e47ff690c6c163290948f80ff97", "host": "lnd_node1_1:9735"}
testrunner_1               | 2022-05-03T20:27:02.960Z	INFO	Open channel	{"node": "node0_0", "peer": "node1_1"}
testrunner_1               | 2022-05-03T20:27:03.069Z	INFO	Connecting	{"node": "node0_0", "peer": "black_hole", "peerPubKey": "030c90b067162358d4e665647674e26c3c716a0fbcb67d2ef3d7f960771db1dc1e", "host": "lnd_black_hole:9735"}
testrunner_1               | 2022-05-03T20:27:03.075Z	INFO	Open channel	{"node": "node0_0", "peer": "black_hole"}
testrunner_1               | 2022-05-03T20:27:03.143Z	INFO	Connecting	{"node": "node0_0", "peer": "black_hole_indirect", "peerPubKey": "02fd6cdbe61b2b5268a846216dea94d3665c79d32d6b851cd05624dd50039972b5", "host": "lnd_black_hole_indirect:9735"}
testrunner_1               | 2022-05-03T20:27:03.150Z	INFO	Open channel	{"node": "node0_0", "peer": "black_hole_indirect"}
testrunner_1               | 2022-05-03T20:27:03.221Z	INFO	Confirm channels
testrunner_1               | 2022-05-03T20:27:03.262Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_3", "expected": 4, "count": 0}
testrunner_1               | 2022-05-03T20:27:04.325Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_3", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.329Z	DEBUG	Waiting for active channels	{"node": "node1_0", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.334Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_4", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.338Z	DEBUG	Waiting for active channels	{"node": "node2_0", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.341Z	DEBUG	Waiting for active channels	{"node": "node0_0", "expected": 5, "count": 5}
testrunner_1               | 2022-05-03T20:27:04.345Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_1", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.349Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_5", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.352Z	DEBUG	Waiting for active channels	{"node": "node2_1", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.357Z	DEBUG	Waiting for active channels	{"node": "destination1", "expected": 8, "count": 8}
testrunner_1               | 2022-05-03T20:27:04.360Z	DEBUG	Waiting for active channels	{"node": "node0_1", "expected": 3, "count": 3}
testrunner_1               | 2022-05-03T20:27:04.362Z	DEBUG	Waiting for active channels	{"node": "node1_1", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.365Z	DEBUG	Waiting for active channels	{"node": "node3_0", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.368Z	DEBUG	Waiting for active channels	{"node": "node4_1", "expected": 3, "count": 3}
testrunner_1               | 2022-05-03T20:27:04.374Z	DEBUG	Waiting for active channels	{"node": "black_hole", "expected": 7, "count": 7}
testrunner_1               | 2022-05-03T20:27:04.377Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_0", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.384Z	DEBUG	Waiting for active channels	{"node": "black_hole_indirect", "expected": 7, "count": 7}
testrunner_1               | 2022-05-03T20:27:04.388Z	DEBUG	Waiting for active channels	{"node": "node4_0", "expected": 3, "count": 3}
testrunner_1               | 2022-05-03T20:27:04.393Z	DEBUG	Waiting for active channels	{"node": "destination2", "expected": 7, "count": 7}
testrunner_1               | 2022-05-03T20:27:04.400Z	DEBUG	Waiting for active channels	{"node": "node3_1", "expected": 5, "count": 5}
testrunner_1               | 2022-05-03T20:27:04.403Z	DEBUG	Waiting for active channels	{"node": "no_liquidity_2", "expected": 4, "count": 4}
testrunner_1               | 2022-05-03T20:27:04.408Z	DEBUG	Waiting for active channels	{"node": "start", "expected": 2, "count": 2}
testrunner_1               | 2022-05-03T20:27:04.408Z	DEBUG	Setting policy	{"node": "node4_0", "channel": "funding_txid_bytes:\"\\r\\253v\\030\\347*`\\002E!\\361\\0355\\326ez\\342n\\275H.\\373\\207\\\"m^\\032\\\"\\364|\\220\\014\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:04.412Z	DEBUG	Setting policy	{"node": "destination1", "channel": "funding_txid_bytes:\"\\r\\253v\\030\\347*`\\002E!\\361\\0355\\326ez\\342n\\275H.\\373\\207\\\"m^\\032\\\"\\364|\\220\\014\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:05.417Z	DEBUG	Setting policy	{"node": "no_liquidity_1", "channel": "funding_txid_bytes:\"\\375U\\270\\004-\\204V\\266Z\\252\\375[\\357\\272:\\226lO\\332'\\223\\243$\\345\\213\\244d!\\241>!\\246\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:05.649Z	DEBUG	Setting policy	{"node": "destination1", "channel": "funding_txid_bytes:\"\\375U\\270\\004-\\204V\\266Z\\252\\375[\\357\\272:\\226lO\\332'\\223\\243$\\345\\213\\244d!\\241>!\\246\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:05.693Z	DEBUG	Setting policy	{"node": "no_liquidity_1", "channel": "funding_txid_bytes:\"\\324\\035\\021<\\342\\202\\020?\\334\\\\\\205\\352g\\033\\017\\n\\233\\212\\211\\344\\3678\\030\\330\\365\\240*x\\317\\227s\\347\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:06.144Z	DEBUG	Setting policy	{"node": "destination2", "channel": "funding_txid_bytes:\"\\324\\035\\021<\\342\\202\\020?\\334\\\\\\205\\352g\\033\\017\\n\\233\\212\\211\\344\\3678\\030\\330\\365\\240*x\\317\\227s\\347\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:06.153Z	DEBUG	Setting policy	{"node": "no_liquidity_3", "channel": "funding_txid_bytes:\"*\\305wh\\304`.@\\030\\252S\\347\\234,\\030\\230r0U\\371\\366\\243\\374\\2056\\331\\225\\354\\017\\264\\276\\232\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:06.695Z	DEBUG	Setting policy	{"node": "destination1", "channel": "funding_txid_bytes:\"*\\305wh\\304`.@\\030\\252S\\347\\234,\\030\\230r0U\\371\\366\\243\\374\\2056\\331\\225\\354\\017\\264\\276\\232\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:06.723Z	DEBUG	Setting policy	{"node": "no_liquidity_3", "channel": "funding_txid_bytes:\"\\023v=Z\\261\\004|Z\\354K\\377\\312US\\227\\335\\337j@7\\343\\361GRV\\214\\247\\214N\\275\\243\\226\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:07.145Z	DEBUG	Setting policy	{"node": "destination2", "channel": "funding_txid_bytes:\"\\023v=Z\\261\\004|Z\\354K\\377\\312US\\227\\335\\337j@7\\343\\361GRV\\214\\247\\214N\\275\\243\\226\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:07.151Z	DEBUG	Setting policy	{"node": "no_liquidity_5", "channel": "funding_txid_bytes:\"\\376\\370\\027Tc\\306?\\224n\\327\\353\\020\\255+\\253\\227\\321\\366_\\322\\247rk\\235\\273\\220F\\364O%\\000U\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:07.720Z	DEBUG	Setting policy	{"node": "destination1", "channel": "funding_txid_bytes:\"\\376\\370\\027Tc\\306?\\224n\\327\\353\\020\\255+\\253\\227\\321\\366_\\322\\247rk\\235\\273\\220F\\364O%\\000U\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:07.724Z	DEBUG	Setting policy	{"node": "no_liquidity_5", "channel": "funding_txid_bytes:\"\\227\\344\\271k\\334\\264\\242_9G\\264:QRKA;\\364<\\221~\\031\\006PS\\310/\\026)\\341%{\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.229Z	DEBUG	Setting policy	{"node": "destination2", "channel": "funding_txid_bytes:\"\\227\\344\\271k\\334\\264\\242_9G\\264:QRKA;\\364<\\221~\\031\\006PS\\310/\\026)\\341%{\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.257Z	DEBUG	Setting policy	{"node": "node0_1", "channel": "funding_txid_bytes:\"\\230h\\362\\257`\\272^F\\005\\342\\224!\\n\\257\\034\\030\\355\\024\\314\\356\\317\\260t\\202\\360&\\317\\022\\255\\005\\247\\233\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.450Z	DEBUG	Setting policy	{"node": "node1_0", "channel": "funding_txid_bytes:\"\\230h\\362\\257`\\272^F\\005\\342\\224!\\n\\257\\034\\030\\355\\024\\314\\356\\317\\260t\\202\\360&\\317\\022\\255\\005\\247\\233\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.474Z	DEBUG	Setting policy	{"node": "node0_1", "channel": "funding_txid_bytes:\"\\372<\\007C\\3060\\261$\\266z\\344d\\013B\\034C\\334\\200QX\\227Y\\273\\317\\315\\205j5C\\231\\343\\245\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.477Z	DEBUG	Setting policy	{"node": "node1_1", "channel": "funding_txid_bytes:\"\\372<\\007C\\3060\\261$\\266z\\344d\\013B\\034C\\334\\200QX\\227Y\\273\\317\\315\\205j5C\\231\\343\\245\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.479Z	DEBUG	Setting policy	{"node": "node3_0", "channel": "funding_txid_bytes:\"\\016f\\200\\235\\321\\013<9\\36267\\243\\373`\\243\\2605]\\333\\350\\322\\034\\255Ie#\\220\\240+\\243?\\027\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.539Z	DEBUG	Setting policy	{"node": "node4_0", "channel": "funding_txid_bytes:\"\\016f\\200\\235\\321\\013<9\\36267\\243\\373`\\243\\2605]\\333\\350\\322\\034\\255Ie#\\220\\240+\\243?\\027\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.735Z	DEBUG	Setting policy	{"node": "node3_0", "channel": "funding_txid_bytes:\"\\373\\325\\344\\205&\\237f\\004\\266\\274jm\\230\\220\\002\\340!\\363\\r\\023:9D\\037\\373C\\\\B\\022\\220\\306\\363\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:08.742Z	DEBUG	Setting policy	{"node": "node4_1", "channel": "funding_txid_bytes:\"\\373\\325\\344\\205&\\237f\\004\\266\\274jm\\230\\220\\002\\340!\\363\\r\\023:9D\\037\\373C\\\\B\\022\\220\\306\\363\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.230Z	DEBUG	Setting policy	{"node": "node3_1", "channel": "funding_txid_bytes:\"\\373\\t\\201,\\264\\276L\\265\\017\\020|\\003\\033W\\257}\\035\\014J\\230\\324\\036\\000F\\252\\321\\356\\366\\rK\\347>\" ", "policy": {"CltvDelta":40,"BaseFee":1000,"FeeRate":10000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.242Z	DEBUG	Setting policy	{"node": "node4_1", "channel": "funding_txid_bytes:\"\\373\\t\\201,\\264\\276L\\265\\017\\020|\\003\\033W\\257}\\035\\014J\\230\\324\\036\\000F\\252\\321\\356\\366\\rK\\347>\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.249Z	DEBUG	Setting policy	{"node": "node3_1", "channel": "funding_txid_bytes:\"\\272\\211hO\\242\\t\\321\\000]3\\325E\\t\\261s\\215\\325h\\325`RdR\\307\\004\\250\\2126z\\220\\260'\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.253Z	DEBUG	Setting policy	{"node": "destination2", "channel": "funding_txid_bytes:\"\\272\\211hO\\242\\t\\321\\000]3\\325E\\t\\261s\\215\\325h\\325`RdR\\307\\004\\250\\2126z\\220\\260'\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.260Z	DEBUG	Setting policy	{"node": "node3_1", "channel": "funding_txid_bytes:\"\\373\\342dF\\033ceQ*\\224Xv\\025\\337\\266\\204R6\\r\\324\\227;$\\265\\262\\260w\\036\\311\\261N\\271\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.264Z	DEBUG	Setting policy	{"node": "node4_0", "channel": "funding_txid_bytes:\"\\373\\342dF\\033ceQ*\\224Xv\\025\\337\\266\\204R6\\r\\324\\227;$\\265\\262\\260w\\036\\311\\261N\\271\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.269Z	DEBUG	Setting policy	{"node": "black_hole", "channel": "funding_txid_bytes:\"\\321$\\333n\\322p\\241\\254|i\\302\\357\\331\\226B\\211+\\004\\200Xa\\324E)\\376\\\"X\\314}\\370\\014\\266\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.272Z	DEBUG	Setting policy	{"node": "no_liquidity_0", "channel": "funding_txid_bytes:\"\\321$\\333n\\322p\\241\\254|i\\302\\357\\331\\226B\\211+\\004\\200Xa\\324E)\\376\\\"X\\314}\\370\\014\\266\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.275Z	DEBUG	Setting policy	{"node": "black_hole", "channel": "funding_txid_bytes:\"\\347'q\\271\\032\\345o\\212\\247\\350\\273\\265\\336\\340 =6u/\\271_z\\351 \\354\\263p\\032\\314H\\365\\377\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.278Z	DEBUG	Setting policy	{"node": "no_liquidity_1", "channel": "funding_txid_bytes:\"\\347'q\\271\\032\\345o\\212\\247\\350\\273\\265\\336\\340 =6u/\\271_z\\351 \\354\\263p\\032\\314H\\365\\377\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.284Z	DEBUG	Setting policy	{"node": "black_hole", "channel": "funding_txid_bytes:\"\\372i\\010\\210\\201&cT\\240\\224j\\220\\016=8\\013\\204\\333\\366\\333W\\370\\025'\\227,\\315\\236O\\322Eh\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.287Z	DEBUG	Setting policy	{"node": "no_liquidity_2", "channel": "funding_txid_bytes:\"\\372i\\010\\210\\201&cT\\240\\224j\\220\\016=8\\013\\204\\333\\366\\333W\\370\\025'\\227,\\315\\236O\\322Eh\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.295Z	DEBUG	Setting policy	{"node": "black_hole", "channel": "funding_txid_bytes:\"Sk\\352J{Od\\032\\202\\027H\\024\\037d\\300>z\\245\\333?\\275gTs\\2537-H9y\\357\\225\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.298Z	DEBUG	Setting policy	{"node": "no_liquidity_3", "channel": "funding_txid_bytes:\"Sk\\352J{Od\\032\\202\\027H\\024\\037d\\300>z\\245\\333?\\275gTs\\2537-H9y\\357\\225\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.311Z	DEBUG	Setting policy	{"node": "black_hole", "channel": "funding_txid_bytes:\"@Z\\200j\\355\\330\\212|T\\351\\254\\341\\247\\225I\\326\\3257)\\333\\253~\\033_\\3529\\352\\346.c\\332\\224\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.320Z	DEBUG	Setting policy	{"node": "no_liquidity_4", "channel": "funding_txid_bytes:\"@Z\\200j\\355\\330\\212|T\\351\\254\\341\\247\\225I\\326\\3257)\\333\\253~\\033_\\3529\\352\\346.c\\332\\224\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.534Z	DEBUG	Setting policy	{"node": "black_hole", "channel": "funding_txid_bytes:\"\\177\\020s~\\301#8\\226\\017\\216\\221\\271\\256\\332\\331\\010\\006\\035\\t\\202r\\037Wm\\236r\\306\\233\\213\\177\\254\\340\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.537Z	DEBUG	Setting policy	{"node": "no_liquidity_5", "channel": "funding_txid_bytes:\"\\177\\020s~\\301#8\\226\\017\\216\\221\\271\\256\\332\\331\\010\\006\\035\\t\\202r\\037Wm\\236r\\306\\233\\213\\177\\254\\340\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.541Z	DEBUG	Setting policy	{"node": "node1_1", "channel": "funding_txid_bytes:\"\\001\\276\\320\\200\\323\\006A\\206M\\246\\374\\204\\342\\320\\235\\242\\236c\\203\\013\\226e\\307I\\255\\263\\247\\222\\244\\250o\\036\" ", "policy": {"CltvDelta":40,"BaseFee":1000,"FeeRate":10000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.543Z	DEBUG	Setting policy	{"node": "node2_0", "channel": "funding_txid_bytes:\"\\001\\276\\320\\200\\323\\006A\\206M\\246\\374\\204\\342\\320\\235\\242\\236c\\203\\013\\226e\\307I\\255\\263\\247\\222\\244\\250o\\036\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:09.997Z	DEBUG	Setting policy	{"node": "node1_1", "channel": "funding_txid_bytes:\"\\014A\\215\\256\\244v\\177h\\031yb\\267\\031oQ!\\240\\275hVb\\352/\\264Xl\\374\\013\\202\\247-\\001\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.003Z	DEBUG	Setting policy	{"node": "node2_1", "channel": "funding_txid_bytes:\"\\014A\\215\\256\\244v\\177h\\031yb\\267\\031oQ!\\240\\275hVb\\352/\\264Xl\\374\\013\\202\\247-\\001\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.007Z	DEBUG	Setting policy	{"node": "node2_1", "channel": "funding_txid_bytes:\"\\364\\2456Q\\257\\307\\243\\211o\\331\\2752h\\264c\\331\\377]y\\305\\000\\303\\342\\375l\\232\\340\\360@\\342\\336\\325\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.009Z	DEBUG	Setting policy	{"node": "node3_1", "channel": "funding_txid_bytes:\"\\364\\2456Q\\257\\307\\243\\211o\\331\\2752h\\264c\\331\\377]y\\305\\000\\303\\342\\375l\\232\\340\\360@\\342\\336\\325\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.012Z	DEBUG	Setting policy	{"node": "node2_1", "channel": "funding_txid_bytes:\"\\237qO\\216\\227\\200I\\314\\0208=t1G\\340\\033\\305e\\014\\276\\302\\313\\327\\037\\030\\332\\327\\177\\337U\\227\\030\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.014Z	DEBUG	Setting policy	{"node": "node3_0", "channel": "funding_txid_bytes:\"\\237qO\\216\\227\\200I\\314\\0208=t1G\\340\\033\\305e\\014\\276\\302\\313\\327\\037\\030\\332\\327\\177\\337U\\227\\030\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.017Z	DEBUG	Setting policy	{"node": "no_liquidity_0", "channel": "funding_txid_bytes:\"Ip\\303\\262\\246\\030\\242E\\376]b\\302+\\ni\\244I \\010\\363\\262\\345\\320k\\275\\t\\002\\355M\\213P\\261\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.020Z	DEBUG	Setting policy	{"node": "destination1", "channel": "funding_txid_bytes:\"Ip\\303\\262\\246\\030\\242E\\376]b\\302+\\ni\\244I \\010\\363\\262\\345\\320k\\275\\t\\002\\355M\\213P\\261\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.024Z	DEBUG	Setting policy	{"node": "no_liquidity_0", "channel": "funding_txid_bytes:\"\\002I\\332S\\212\\013\\372i\\325n\\215\\200n\\372\\006#\\203\\257C\\335\\353\\241\\272\\224\\214\\320\\256\\016Z`\\376\\235\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.026Z	DEBUG	Setting policy	{"node": "destination2", "channel": "funding_txid_bytes:\"\\002I\\332S\\212\\013\\372i\\325n\\215\\200n\\372\\006#\\203\\257C\\335\\353\\241\\272\\224\\214\\320\\256\\016Z`\\376\\235\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.033Z	DEBUG	Setting policy	{"node": "no_liquidity_2", "channel": "funding_txid_bytes:\"\\261W\\002\\341\\316u\\035\\236\\340\\n\\212\\036\\300Na(\\217\\264\\321\\t\\347\\025}9\\352\\371\\303\\252\\006\\342w\\216\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.038Z	DEBUG	Setting policy	{"node": "destination1", "channel": "funding_txid_bytes:\"\\261W\\002\\341\\316u\\035\\236\\340\\n\\212\\036\\300Na(\\217\\264\\321\\t\\347\\025}9\\352\\371\\303\\252\\006\\342w\\216\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.041Z	DEBUG	Setting policy	{"node": "no_liquidity_2", "channel": "funding_txid_bytes:\"\\244\\266\\02733@\\204\\346\\242c\\207~R\\347u\\324\\364\\033\\340[\\021A\\261\\342Tz3S)\\215\\351\\030\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.045Z	DEBUG	Setting policy	{"node": "destination2", "channel": "funding_txid_bytes:\"\\244\\266\\02733@\\204\\346\\242c\\207~R\\347u\\324\\364\\033\\340[\\021A\\261\\342Tz3S)\\215\\351\\030\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.058Z	DEBUG	Setting policy	{"node": "start", "channel": "funding_txid_bytes:\"\\360#\\022\\333\\311\\234?\\354?\\017\\204\\267\\316\\252\\202\\320\\230\\314u1\\006\\265\\264\\354\\256\\372 \\232:0v\\303\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.069Z	DEBUG	Setting policy	{"node": "node0_0", "channel": "funding_txid_bytes:\"\\360#\\022\\333\\311\\234?\\354?\\017\\204\\267\\316\\252\\202\\320\\230\\314u1\\006\\265\\264\\354\\256\\372 \\232:0v\\303\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.073Z	DEBUG	Setting policy	{"node": "start", "channel": "funding_txid_bytes:\"a\\355\\3156#\\035\\224\\2509\\266S\\370\\276U\\027\\356\\207Q\\250\\177e\\3151\\232\\347\\365\\222W\\035\\301Jh\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.076Z	DEBUG	Setting policy	{"node": "node0_1", "channel": "funding_txid_bytes:\"a\\355\\3156#\\035\\224\\2509\\266S\\370\\276U\\027\\356\\207Q\\250\\177e\\3151\\232\\347\\365\\222W\\035\\301Jh\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.080Z	DEBUG	Setting policy	{"node": "node1_0", "channel": "funding_txid_bytes:\"\\250\\354\\210,\\264\\301?f\\026\\363\\300tn\\024\\221\\270\\201 \\334<\\2618\\222\\222\\210\\331X\\026\\013P\\243\\247\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.087Z	DEBUG	Setting policy	{"node": "node2_0", "channel": "funding_txid_bytes:\"\\250\\354\\210,\\264\\301?f\\026\\363\\300tn\\024\\221\\270\\201 \\334<\\2618\\222\\222\\210\\331X\\026\\013P\\243\\247\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.094Z	DEBUG	Setting policy	{"node": "node1_0", "channel": "funding_txid_bytes:\"\\002]\\353\\033t\\243\\221\\260\\207\\0006O\\207\\243\\350]\\2442>D\\2756g\\335\\302\\350\\372\\230\\324{H\\217\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.099Z	DEBUG	Setting policy	{"node": "node2_1", "channel": "funding_txid_bytes:\"\\002]\\353\\033t\\243\\221\\260\\207\\0006O\\207\\243\\350]\\2442>D\\2756g\\335\\302\\350\\372\\230\\324{H\\217\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.104Z	DEBUG	Setting policy	{"node": "node2_0", "channel": "funding_txid_bytes:\"S+G\\250\\372\\223\\370\\217\\207\\024<B\\026\\006\\222\\3004\\260|;\\302\\000%\\0025\\310\\322\\276\\274\\334z\\221\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.107Z	DEBUG	Setting policy	{"node": "node3_0", "channel": "funding_txid_bytes:\"S+G\\250\\372\\223\\370\\217\\207\\024<B\\026\\006\\222\\3004\\260|;\\302\\000%\\0025\\310\\322\\276\\274\\334z\\221\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.110Z	DEBUG	Setting policy	{"node": "node2_0", "channel": "funding_txid_bytes:\"\\322Rw/5\\037\\335\\035\\240)\\276\\215\\017\\001p|\\253|~\\336\\032\\250\\016?]\\332o\\035-\\313\\321Q\" ", "policy": {"CltvDelta":40,"BaseFee":1000,"FeeRate":10000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.113Z	DEBUG	Setting policy	{"node": "node3_1", "channel": "funding_txid_bytes:\"\\322Rw/5\\037\\335\\035\\240)\\276\\215\\017\\001p|\\253|~\\336\\032\\250\\016?]\\332o\\035-\\313\\321Q\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.117Z	DEBUG	Setting policy	{"node": "node4_1", "channel": "funding_txid_bytes:\"7\\363spt\\365\\222\\331\\243>>\\337\\001j\\336']\\337^^\\321\\354\\025\\231$_\\305;5\\035+\\025\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.123Z	DEBUG	Setting policy	{"node": "destination1", "channel": "funding_txid_bytes:\"7\\363spt\\365\\222\\331\\243>>\\337\\001j\\336']\\337^^\\321\\354\\025\\231$_\\305;5\\035+\\025\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.125Z	DEBUG	Setting policy	{"node": "black_hole_indirect", "channel": "funding_txid_bytes:\"g\\253;\\254H\\270\\243v\\\\\\021\\262\\262D7\\315\\243S\\216 \\376\\334R\\236\\220#\\247\\223\\3635)\\335x\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.129Z	DEBUG	Setting policy	{"node": "no_liquidity_3", "channel": "funding_txid_bytes:\"g\\253;\\254H\\270\\243v\\\\\\021\\262\\262D7\\315\\243S\\216 \\376\\334R\\236\\220#\\247\\223\\3635)\\335x\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.346Z	DEBUG	Setting policy	{"node": "black_hole_indirect", "channel": "funding_txid_bytes:\"\\3544\\251Ea\\360\\360;e_\\321H\\331\\356S\\010\\256\\354'$\\306\\253Zs\\025\\366\\200\\256\\230u*\\\"\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.352Z	DEBUG	Setting policy	{"node": "no_liquidity_4", "channel": "funding_txid_bytes:\"\\3544\\251Ea\\360\\360;e_\\321H\\331\\356S\\010\\256\\354'$\\306\\253Zs\\025\\366\\200\\256\\230u*\\\"\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.356Z	DEBUG	Setting policy	{"node": "black_hole_indirect", "channel": "funding_txid_bytes:\"\\034\\037d\\032\\251\\234\\260\\256\\320\\024jj\\212ZO\\276\\312\\177/\\377<\\2507\\262\\002\\365\\354\\002A\\346\\031!\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.361Z	DEBUG	Setting policy	{"node": "no_liquidity_5", "channel": "funding_txid_bytes:\"\\034\\037d\\032\\251\\234\\260\\256\\320\\024jj\\212ZO\\276\\312\\177/\\377<\\2507\\262\\002\\365\\354\\002A\\346\\031!\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.366Z	DEBUG	Setting policy	{"node": "black_hole_indirect", "channel": "funding_txid_bytes:\"\\372\\334\\243o\\207\\232R\\244\\177\\207+\\373In\\\"rk\\266\\221\\325\\356\\256\\037\\014\\037\\263\\362\\242\\326\\354\\210r\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.370Z	DEBUG	Setting policy	{"node": "no_liquidity_0", "channel": "funding_txid_bytes:\"\\372\\334\\243o\\207\\232R\\244\\177\\207+\\373In\\\"rk\\266\\221\\325\\356\\256\\037\\014\\037\\263\\362\\242\\326\\354\\210r\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.377Z	DEBUG	Setting policy	{"node": "black_hole_indirect", "channel": "funding_txid_bytes:\"\\223\\367\\004\\352F\\006r\\317\\331^X\\313\\262\\262\\245\\330\\353\\2302P\\267\\202`N\\007x\\372\\2479\\237\\020\\016\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.382Z	DEBUG	Setting policy	{"node": "no_liquidity_1", "channel": "funding_txid_bytes:\"\\223\\367\\004\\352F\\006r\\317\\331^X\\313\\262\\262\\245\\330\\353\\2302P\\267\\202`N\\007x\\372\\2479\\237\\020\\016\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.389Z	DEBUG	Setting policy	{"node": "black_hole_indirect", "channel": "funding_txid_bytes:\"\\325\\021o\\313M\\250b\\265jCPO\\207{\\216-\\017\\3700\\264\\204\\343\\026k\\367\\306\\267\\021\\221\\316\\323\\232\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.393Z	DEBUG	Setting policy	{"node": "no_liquidity_2", "channel": "funding_txid_bytes:\"\\325\\021o\\313M\\250b\\265jCPO\\207{\\216-\\017\\3700\\264\\204\\343\\026k\\367\\306\\267\\021\\221\\316\\323\\232\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.396Z	DEBUG	Setting policy	{"node": "no_liquidity_4", "channel": "funding_txid_bytes:\"\\247{J\\326\\345\\003\\246\\356Wq\\237\\2737\\264)\\022o\\247)\\r'\\255(O\\325\\030?\\247\\344\\320\\036\\231\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.398Z	DEBUG	Setting policy	{"node": "destination1", "channel": "funding_txid_bytes:\"\\247{J\\326\\345\\003\\246\\356Wq\\237\\2737\\264)\\022o\\247)\\r'\\255(O\\325\\030?\\247\\344\\320\\036\\231\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.401Z	DEBUG	Setting policy	{"node": "no_liquidity_4", "channel": "funding_txid_bytes:\"\\034\\266\\356\\364}zu\\231}\\272\\036\\267\\315\\271\\207\\245\\315\\264\\351\\335\\343\\275\\361\\001\\315\\212:\\263\\225\\306\\266$\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.405Z	DEBUG	Setting policy	{"node": "destination2", "channel": "funding_txid_bytes:\"\\034\\266\\356\\364}zu\\231}\\272\\036\\267\\315\\271\\207\\245\\315\\264\\351\\335\\343\\275\\361\\001\\315\\212:\\263\\225\\306\\266$\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.639Z	DEBUG	Setting policy	{"node": "node0_0", "channel": "funding_txid_bytes:\"T7@\\336\\010\\312<\\254\\262E\\235\\225\\030\\326\\177\\nT\\266\\262*\\312k\\235\\205R\\033\\3528\\013\\017R\\346\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.701Z	DEBUG	Setting policy	{"node": "node1_0", "channel": "funding_txid_bytes:\"T7@\\336\\010\\312<\\254\\262E\\235\\225\\030\\326\\177\\nT\\266\\262*\\312k\\235\\205R\\033\\3528\\013\\017R\\346\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.704Z	DEBUG	Setting policy	{"node": "node0_0", "channel": "funding_txid_bytes:\"\\036\\3779\\323\\221\\354\\244}\\320\\220]\\327\\202\\036\\221\\257k\\031\\031W\\256\\006\\037\\304\\320g\\266,G\\322D\\317\" ", "policy": {"CltvDelta":40,"BaseFee":1000,"FeeRate":10000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.708Z	DEBUG	Setting policy	{"node": "node1_1", "channel": "funding_txid_bytes:\"\\036\\3779\\323\\221\\354\\244}\\320\\220]\\327\\202\\036\\221\\257k\\031\\031W\\256\\006\\037\\304\\320g\\266,G\\322D\\317\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.712Z	DEBUG	Setting policy	{"node": "node0_0", "channel": "funding_txid_bytes:\"*\\333\\355\\340w2w\\241P_}\\r\\030r\\227\\3011C\\321\\217\\333\\217\\270\\314\\365l\\236\\212f\\347\\016\\212\" ", "policy": {"CltvDelta":40,"BaseFee":0,"FeeRate":0,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.714Z	DEBUG	Setting policy	{"node": "black_hole", "channel": "funding_txid_bytes:\"*\\333\\355\\340w2w\\241P_}\\r\\030r\\227\\3011C\\321\\217\\333\\217\\270\\314\\365l\\236\\212f\\347\\016\\212\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.718Z	DEBUG	Setting policy	{"node": "node0_0", "channel": "funding_txid_bytes:\"<)\\205]\\020\\251/\\026-\\357[\\247\\217\\264\\254(\\211\\223b\\241\\222\\301\\312z\\372\\307\\243\\306\\206\\325\\270\\371\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.720Z	DEBUG	Setting policy	{"node": "black_hole_indirect", "channel": "funding_txid_bytes:\"<)\\205]\\020\\251/\\026-\\357[\\247\\217\\264\\254(\\211\\223b\\241\\222\\301\\312z\\372\\307\\243\\306\\206\\325\\270\\371\" ", "policy": {"CltvDelta":40,"BaseFee":100,"FeeRate":1000,"HtlcMaxSat":100000}}
testrunner_1               | 2022-05-03T20:27:10.730Z	INFO	Wait for propagation
testrunner_1               | 2022-05-03T20:27:20.733Z	INFO	Sending payment	{"invoice": "lnbcrt500u1p38ryq7pp52van3p76fxz33ce6jlhjheygvy4gn32fc7t5pwexeequfs4x8sdqdqqcqzpgxqyz5vqsp5vg67xw9rga4wzgcugxyve3pex99dgzykkr7ualdrewe02xqwktfq9qyyssq4vp5sme84kfma6exaus2g9rtskr4xmv75kxfss0jrkhpdvpkkvjjxexf2naygqcew4p6fukylhk6mm8ucl5xc40us309ygn4j3du5jqqcs6s2j"}
testrunner_1               | 2022-05-03T20:27:20.794Z	DEBUG	Payment update	{"htlcIdx": 0, "route": ["node0_0", "black_hole", "no_liquidity_5", "destination1"]}
testrunner_1               | 2022-05-03T20:27:21.047Z	DEBUG	Payment update	{"htlcIdx": 1, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination1"]}
testrunner_1               | 2022-05-03T20:27:21.308Z	DEBUG	Payment update	{"htlcIdx": 2, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination1"]}
testrunner_1               | 2022-05-03T20:27:21.561Z	DEBUG	Payment update	{"htlcIdx": 3, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:21.931Z	DEBUG	Payment update	{"htlcIdx": 4, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination1"]}
testrunner_1               | 2022-05-03T20:27:22.194Z	DEBUG	Payment update	{"htlcIdx": 5, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination1"]}
testrunner_1               | 2022-05-03T20:27:22.441Z	DEBUG	Payment update	{"htlcIdx": 6, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination1"]}
testrunner_1               | 2022-05-03T20:27:22.905Z	DEBUG	Payment update	{"htlcIdx": 7, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:23.330Z	DEBUG	Payment update	{"htlcIdx": 8, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination1"]}
testrunner_1               | 2022-05-03T20:27:23.791Z	DEBUG	Payment update	{"htlcIdx": 9, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination1"]}
testrunner_1               | 2022-05-03T20:27:24.153Z	DEBUG	Payment update	{"htlcIdx": 10, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination1"]}
testrunner_1               | 2022-05-03T20:27:24.519Z	DEBUG	Payment update	{"htlcIdx": 11, "route": ["node0_0", "black_hole", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:24.882Z	DEBUG	Payment update	{"htlcIdx": 12, "route": ["node0_0", "node1_0", "node2_1", "node3_1", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:25.020Z	DEBUG	Payment update	{"htlcIdx": 13, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:25.412Z	DEBUG	Payment update	{"htlcIdx": 14, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:25.792Z	DEBUG	Payment update	{"htlcIdx": 15, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:26.148Z	DEBUG	Payment update	{"htlcIdx": 16, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:26.497Z	DEBUG	Payment update	{"htlcIdx": 17, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:27.020Z	DEBUG	Payment update	{"htlcIdx": 18, "route": ["node0_0", "node1_1", "node2_1", "node3_1", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:27.271Z	DEBUG	Payment update	{"htlcIdx": 19, "route": ["node0_0", "node1_1", "node2_0", "node3_0", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:27.721Z	DEBUG	Payment update	{"htlcIdx": 20, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination1"]}
testrunner_1               | 2022-05-03T20:27:27.736Z	DEBUG	Payment update	{"htlcIdx": 21, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination1"]}
testrunner_1               | 2022-05-03T20:27:28.112Z	DEBUG	Payment update	{"htlcIdx": 22, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination1"]}
testrunner_1               | 2022-05-03T20:27:28.142Z	DEBUG	Payment update	{"htlcIdx": 23, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination1"]}
testrunner_1               | 2022-05-03T20:27:28.499Z	DEBUG	Payment update	{"htlcIdx": 24, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination1"]}
testrunner_1               | 2022-05-03T20:27:28.529Z	DEBUG	Payment update	{"htlcIdx": 25, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:29.058Z	DEBUG	Payment update	{"htlcIdx": 26, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination1"]}
testrunner_1               | 2022-05-03T20:27:29.101Z	DEBUG	Payment update	{"htlcIdx": 27, "route": ["node0_0", "black_hole", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:29.393Z	DEBUG	Payment update	{"htlcIdx": 28, "route": ["node0_0", "node1_1", "node2_1", "node3_0", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:29.426Z	DEBUG	Payment update	{"htlcIdx": 29, "route": ["node0_0", "node1_1", "node2_1", "node3_1", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:29.709Z	DEBUG	Payment update	{"htlcIdx": 30, "route": ["node0_0", "node1_1", "node2_0", "node3_0", "node4_1", "destination1"]}
testrunner_1               | 2022-05-03T20:27:29.763Z	DEBUG	Payment update	{"htlcIdx": 31, "route": ["node0_0", "node1_1", "node2_0", "node3_0", "node4_1", "destination1"]}
testrunner_1               | 2022-05-03T20:27:30.035Z	DEBUG	Payment update	{"htlcIdx": 32, "route": ["node0_0", "node1_0", "node2_1", "node3_1", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:30.090Z	DEBUG	Payment update	{"htlcIdx": 33, "route": ["node0_0", "node1_0", "node2_1", "node3_1", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:30.151Z	DEBUG	Payment update	{"htlcIdx": 34, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:30.207Z	DEBUG	Payment update	{"htlcIdx": 35, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:30.500Z	DEBUG	Payment update	{"htlcIdx": 36, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:30.555Z	DEBUG	Payment update	{"htlcIdx": 37, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:30.912Z	DEBUG	Payment update	{"htlcIdx": 38, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:30.966Z	DEBUG	Payment update	{"htlcIdx": 39, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:31.242Z	DEBUG	Payment update	{"htlcIdx": 40, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:31.311Z	DEBUG	Payment update	{"htlcIdx": 41, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:31.634Z	DEBUG	Payment update	{"htlcIdx": 42, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:31.760Z	DEBUG	Payment update	{"htlcIdx": 43, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2", "no_liquidity_4", "destination1"]}
testrunner_1               | 2022-05-03T20:27:32.015Z	DEBUG	Payment update	{"htlcIdx": 44, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:32.090Z	DEBUG	Payment update	{"htlcIdx": 45, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_0", "destination1"]}
testrunner_1               | 2022-05-03T20:27:32.478Z	DEBUG	Payment update	{"htlcIdx": 46, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_1", "destination1"]}
testrunner_1               | 2022-05-03T20:27:32.537Z	DEBUG	Payment update	{"htlcIdx": 47, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "node4_1", "destination1"]}
testrunner_1               | 2022-05-03T20:27:33.305Z	INFO	Time elapsed	{"time": "12.571239173s"}
testrunner_1               | 2022-05-03T20:27:33.305Z	INFO	Sending payment	{"invoice": "lnbcrt500u1p38ryq7pp53a7h6lsqnefp98hjnmzmlplcwz25xmhf9qkyccr7v9ryhupd4s5qdqqcqzpgxqyz5vqsp57xledw9stpgus9s4wehg6jaxkdd2chwvxft080fullqznh5g6urq9qyyssqgzs38tn3zcc4pkkmwcaul2d8htd0v5jazwzkmhzh9h8tc8vj762qyrkk2ze39rf888zal3czd3m5qh9u8mt3ezymwkumrlgavaatdyqpjkveww"}
testrunner_1               | 2022-05-03T20:27:33.335Z	DEBUG	Payment update	{"htlcIdx": 0, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2"]}
testrunner_1               | 2022-05-03T20:27:33.355Z	DEBUG	Payment update	{"htlcIdx": 1, "route": ["node0_0", "black_hole_indirect", "no_liquidity_5", "destination2"]}
testrunner_1               | 2022-05-03T20:27:33.826Z	DEBUG	Payment update	{"htlcIdx": 2, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2"]}
testrunner_1               | 2022-05-03T20:27:33.858Z	DEBUG	Payment update	{"htlcIdx": 3, "route": ["node0_0", "black_hole_indirect", "no_liquidity_2", "destination2"]}
testrunner_1               | 2022-05-03T20:27:34.162Z	DEBUG	Payment update	{"htlcIdx": 4, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2"]}
testrunner_1               | 2022-05-03T20:27:34.206Z	DEBUG	Payment update	{"htlcIdx": 5, "route": ["node0_0", "black_hole_indirect", "no_liquidity_3", "destination2"]}
testrunner_1               | 2022-05-03T20:27:34.508Z	DEBUG	Payment update	{"htlcIdx": 6, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2"]}
testrunner_1               | 2022-05-03T20:27:34.555Z	DEBUG	Payment update	{"htlcIdx": 7, "route": ["node0_0", "black_hole_indirect", "no_liquidity_0", "destination2"]}
testrunner_1               | 2022-05-03T20:27:34.867Z	DEBUG	Payment update	{"htlcIdx": 8, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2"]}
testrunner_1               | 2022-05-03T20:27:34.920Z	DEBUG	Payment update	{"htlcIdx": 9, "route": ["node0_0", "black_hole_indirect", "no_liquidity_1", "destination2"]}
testrunner_1               | 2022-05-03T20:27:35.193Z	DEBUG	Payment update	{"htlcIdx": 10, "route": ["node0_0", "node1_0", "node2_1", "node3_1", "destination2"]}
testrunner_1               | 2022-05-03T20:27:35.249Z	DEBUG	Payment update	{"htlcIdx": 11, "route": ["node0_0", "node1_0", "node2_1", "node3_1", "destination2"]}
testrunner_1               | 2022-05-03T20:27:35.314Z	DEBUG	Payment update	{"htlcIdx": 12, "route": ["node0_0", "node1_1", "node2_1", "node3_1", "destination2"]}
testrunner_1               | 2022-05-03T20:27:35.382Z	DEBUG	Payment update	{"htlcIdx": 13, "route": ["node0_0", "node1_1", "node2_1", "node3_1", "destination2"]}
testrunner_1               | 2022-05-03T20:27:35.547Z	DEBUG	Payment update	{"htlcIdx": 14, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "destination2"]}
testrunner_1               | 2022-05-03T20:27:35.691Z	DEBUG	Payment update	{"htlcIdx": 15, "route": ["node0_0", "node1_1", "node2_0", "node3_1", "destination2"]}
testrunner_1               | 2022-05-03T20:27:36.232Z	INFO	Time elapsed	{"time": "2.927482501s"}
testrunner_1               | 2022-05-03T20:27:36.233Z	INFO	Total time elapsed	{"time": "15.49923859s"}
testrunner_1               | 2022-05-03T20:27:36.233Z	INFO	Done
```