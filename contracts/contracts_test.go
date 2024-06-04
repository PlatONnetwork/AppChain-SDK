package contracts

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/test-go/testify/assert"
	"testing"
)

func TestMapG1(t *testing.T) {
	input1 := common.Hex2Bytes("0000000000000000000000000000000009cf51f5816625ccdae5e5fdf1af8005d78b7af4e38d7bbb8f92d60a09453b8cdb1882259e0b5ed089a92d925e639158")
	c := &bls12381MapG1{}
	output1, err := c.Run(input1)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("0x%x", output1))

	input2 := common.Hex2Bytes("000000000000000000000000000000000a1e63e25d588a78d647492ccc543388344100664f413efa809d39e8b060b762d1e76c5795ebfefd17695f52df5ac78b")
	output2, err := c.Run(input2)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("0x%x", output2))

	input3 := common.Hex2Bytes("000000000000000000000000000000000cddf2a69c25ded51b28453dd92061d04b213bb73e42d7ab13ef6a491b5b6ce9248dd2cf5128ef3ebd37ff6797f565d5")
	output3, err := c.Run(input3)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("0x%x", output3))

	c1 := &bls12381G1Add{}
	pub, err := c1.Run(append(output1, output2...))
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("0x%x", pub))
	aggpub, err := c1.Run(append(pub, output3...))
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("agg: 0x%x", aggpub))

	c3 := &bls12381MapG2{}
	input4 := common.Hex2Bytes("000000000000000000000000000000000fc3ef66c713d08c0bef2cd21e87f8f633a558fa443294f00cf8781dbc0ac3374c82877b219e722e4a9251199a02edb60000000000000000000000000000000013bd816f0c3c8c54912c3df365ffa7b86369be3c47cf68205b6be5c869b592ab8f35b8af4052972ac47c47a2583e5692")
	sig, err := c3.Run(input4)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("sig: 0x%x", sig))

	input5 := common.Hex2Bytes("000000000000000000000000000000000202be61f38066dea35415c6c77c97ad133d4461d8cf740b087e15285d34bc777e1495127740a20b3408c2ee0be7343e000000000000000000000000000000000ef99eaa973a8a54e95a039683e0d8543899c5db785bfdac01a5e1f62f9a6b22ecee3575cf04a6ed8dd4866a6b11cb8b")
	fp, err := c3.Run(input5)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("fp: 0x%x", fp))

	input6 := common.Hex2Bytes("0000000000000000000000000000000006da95e052941571f87ae88c9513599026939b49d86ce7cdb731cf34306d3c797a5077456a7136d78998279cf235ee69000000000000000000000000000000000acb2831515ba0b4ab0eeeb137fd02668ac26a2f79c58ef8bbba093434532e6d1206d5027794afd91fc3dce8531b5867")
	sp, err := c3.Run(input6)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("sp: 0x%x", sp))

	c4 := &bls12381G2Add{}
	input7 := common.Hex2Bytes("000000000000000000000000000000000aa0f7cf0d993c383984caaa6ef72603232c6e156a055b825f954edede70a191fcb2dda859cc6f25bc30be322629ee6d00000000000000000000000000000000054ce6824b303cccae8b3ea058fec360509933ec36cbe80637951dea12ad8ce181fbc816003473e15a69ec7a788c1442000000000000000000000000000000000270e544ac2d13f782a811a63969be076e98cd1c63fba35ee753558b66c6ac0e7425b2c4159198b2e553316d9d3263f00000000000000000000000000000000009435339be48282a857e6dd6e9219f3e4a4126ce8b1d1ed4c7635ab06e7d911bded586e745c5926ca47b21ab22e4a4bb0000000000000000000000000000000003a1c904401f3899945952f7e629db32d952c55c06afa1182e63d034cc6950f0c60e0aa89df4891622e402b41e0745070000000000000000000000000000000005e9866c3d46b3878a7ec4e46e17c11886d00ea8caf4cfa79b149f07ffdef8331c2b1c30f0c66a80b605f383c6377cd3000000000000000000000000000000000c6928b156792f99b639a7c077f4700334c2a2466ca9b0304974afcd2b13ee49b9cb606b3a075b1403ac99137076e41600000000000000000000000000000000012e25450b9717e80def375dbb791099caa080d5b4e80756cfde72e526debac5d7352fa8685e94de2806c448ca416410")
	hashCurve, err := c4.Run(input7)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("hash: 0x%x", hashCurve))
}

func TestVerifyBls(t *testing.T) {
	c := &bls12381Pairing{}
	input := common.Hex2Bytes("00000000000000000000000000000000186494199e8eb845d7837f4871e60dacef4febf5ffd9e33d6159f201cbdfc1dad1f127c4b55044169193afb51d9995860000000000000000000000000000000001940e16efd26b2891fe84295453a781aee8767a57e4de9e26e26981cfe94febee62910a1b099943a104226581023f9100000000000000000000000000000000118ab0e5c44538417dc1ccc6c74cf63a81b00605ecabb8c50e39c78c6feb243cdeea1f7cc9069788a3972a0e2ccbff090000000000000000000000000000000005cea222f88a75a3eaf9df3a55d6adf73730f61ab86612eb7f32ad3477b784209dab8fc9a4976679c2d62df70e21489a00000000000000000000000000000000198546c2699280905bea44a8c3dc11471b6ce654a491e54737e469ecd2ae25dccabe2c0995a9b510738a5cc6648da9a00000000000000000000000000000000003b2f71bc6f5f30eedcf01166fc7ecd9dd355a560d1aea9c38f57276eaab10495d23f2c8666591e0a2c52dbe0a52a55c0000000000000000000000000000000017f1d3a73197d7942695638c4fa9ac0fc3688c4f9774b905a14e3a3f171bac586c55e83ff97a1aeffb3af00adb22c6bb00000000000000000000000000000000114d1d6855d545a8aa7d76c8cf2e21f267816aef1db507c96655b9d5caac42364e6f38ba0ecb751bad54dcd6b939c2ca000000000000000000000000000000000b87c14f6ad8b19edc0a4d7b4743f71db5bcb33c2e06a894cb11b7dba11d55d1360a76268429b68e382693653b99e84b000000000000000000000000000000000ffc47794f633d29e15a014c841fa33c9ba0070e95b2c059d550477192556d90dc1d746dbb6e404f777af8a49d645f510000000000000000000000000000000002da1674d16a3146d46d2637397cfc59ec6acd5c1a948ba124e81b93326b2c92c20dd61f2e7ed9c902f325aa998fe925000000000000000000000000000000000596fa3e213fe450754a0b7900c451672a8293ce7abecab5d112da6ba3a3b3c3fbe176b4c81ad4504c394bc178717272")
	output, err := c.Run(input)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("%x", output))
}

func TestVerify(t *testing.T) {
	var aggpub bls.PublicKey

	b1 := common.Hex2Bytes("89cf51f5816625ccdae5e5fdf1af8005d78b7af4e38d7bbb8f92d60a09453b8cdb1882259e0b5ed089a92d925e639158")
	b2 := common.Hex2Bytes("8a1e63e25d588a78d647492ccc543388344100664f413efa809d39e8b060b762d1e76c5795ebfefd17695f52df5ac78b")
	b3 := common.Hex2Bytes("acddf2a69c25ded51b28453dd92061d04b213bb73e42d7ab13ef6a491b5b6ce9248dd2cf5128ef3ebd37ff6797f565d5")

	assert.NoError(t, aggpub.Deserialize(b1))
	fmt.Println(fmt.Sprintf("0x%x", aggpub.SerializeUncompressed()))
	var p2 bls.PublicKey
	assert.NoError(t, p2.Deserialize(b2))
	var p3 bls.PublicKey
	assert.NoError(t, p3.Deserialize(b3))

	aggpub.Add(&p2)
	aggpub.Add(&p3)

	ab := aggpub.Serialize()
	fmt.Println(fmt.Sprintf("0x%x", ab))
	fmt.Println(fmt.Sprintf("0x%x", aggpub.SerializeUncompressed()))

	sb := common.Hex2Bytes("93bd816f0c3c8c54912c3df365ffa7b86369be3c47cf68205b6be5c869b592ab8f35b8af4052972ac47c47a2583e56920fc3ef66c713d08c0bef2cd21e87f8f633a558fa443294f00cf8781dbc0ac3374c82877b219e722e4a9251199a02edb6")
	var sig bls.Sign
	assert.NoError(t, sig.Deserialize(sb))

	hash := common.Hex2Bytes("c0bc3c9414679391e12348bcd5af1723b3d2d1dea93bd724637f028a5b50b50b")

	assert.True(t, sig.Verify(&aggpub, string(hash)))

	var aggpub1 bls.PublicKey
	assert.NoError(t, aggpub1.Deserialize(common.Hex2Bytes("986494199e8eb845d7837f4871e60dacef4febf5ffd9e33d6159f201cbdfc1dad1f127c4b55044169193afb51d999586")))
	assert.True(t, sig.Verify(&aggpub1, string(hash)))

}

func TestMapG1Cmp(t *testing.T) {
	b1 := common.Hex2Bytes("89cf51f5816625ccdae5e5fdf1af8005d78b7af4e38d7bbb8f92d60a09453b8cdb1882259e0b5ed089a92d925e639158")
	b1[0] = b1[0] & 0x1f
	var input [64]byte
	diff := 64 - len(b1)
	copy(input[diff:], b1)
	fmt.Println(fmt.Sprintf("0x%x", input))
	c := &bls12381MapG1{}
	g1, err := c.Run(input[:])
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("p1: 0x%x", g1))

	b2 := common.Hex2Bytes("8a1e63e25d588a78d647492ccc543388344100664f413efa809d39e8b060b762d1e76c5795ebfefd17695f52df5ac78b")
	b2[0] = b2[0] & 0x1f
	copy(input[diff:], b2)
	g2, err := c.Run(input[:])
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("p2: 0x%x", g2))

	b3 := common.Hex2Bytes("acddf2a69c25ded51b28453dd92061d04b213bb73e42d7ab13ef6a491b5b6ce9248dd2cf5128ef3ebd37ff6797f565d5")
	b3[0] = b3[0] & 0x1f
	copy(input[diff:], b3)
	g3, err := c.Run(input[:])
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("p2: 0x%x", g3))

	gadd := &bls12381G1Add{}
	g12, err := gadd.Run(append(g1, g2...))
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("g12: 0x%x", g12))

	agg, err := gadd.Run(append(g12, g3...))
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("agg: 0x%x", agg))
	fmt.Println(len(agg))

	aggbytes := common.Hex2Bytes("186494199e8eb845d7837f4871e60dacef4febf5ffd9e33d6159f201cbdfc1dad1f127c4b55044169193afb51d99958601940e16efd26b2891fe84295453a781aee8767a57e4de9e26e26981cfe94febee62910a1b099943a104226581023f91")

	var aggPub bls.PublicKey
	err = aggPub.DeserializeUncompressed(aggbytes)
	assert.NoError(t, err)
	fmt.Println(fmt.Sprintf("aggPub: 0x%x", aggPub.Serialize()))
}

func TestSerialize(t *testing.T) {
	var pub bls.PublicKey
	b := common.Hex2Bytes("1876dd03316ff007a2efb4c5f452d8418edacc2881b20e8340895f6fc768d14fd89bd9db3dcfb53fa98a1e96055fa83e")
	b[0] = b[0] | 0x1f
	assert.NoError(t, pub.Deserialize(b))
}

func TestUncompressed(t *testing.T) {
	b1 := common.Hex2Bytes("89cf51f5816625ccdae5e5fdf1af8005d78b7af4e38d7bbb8f92d60a09453b8cdb1882259e0b5ed089a92d925e639158")
	b2 := common.Hex2Bytes("8a1e63e25d588a78d647492ccc543388344100664f413efa809d39e8b060b762d1e76c5795ebfefd17695f52df5ac78b")
	b3 := common.Hex2Bytes("acddf2a69c25ded51b28453dd92061d04b213bb73e42d7ab13ef6a491b5b6ce9248dd2cf5128ef3ebd37ff6797f565d5")

	var p1 bls.PublicKey
	p1.Deserialize(b1)
	fmt.Println(fmt.Sprintf("%x", p1.SerializeUncompressed()))

	var p2 bls.PublicKey
	p2.Deserialize(b2)
	fmt.Println(fmt.Sprintf("%x", p2.SerializeUncompressed()))

	var p3 bls.PublicKey
	p3.Deserialize(b3)
	fmt.Println(fmt.Sprintf("%x", p3.SerializeUncompressed()))
}
