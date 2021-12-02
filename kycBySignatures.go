package circuits

import (
	"errors"
	"math/big"
	"strconv"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-merkletree-sql"
)

const (
	// KycBySignaturesPublicSignalsSchema is schema to parse json data for additional information
	KycBySignaturesPublicSignalsSchema PublicSchemaJSON = `{"user_identifier":0,"challenge":1,"countryClaimIssuerId":2,"countryClaimIssuerBBJIdenState":3,"countryBlacklist_1":4,"countryBlacklist_2":5,"countryBlacklist_3":6,"countryBlacklist_4":7,"countryBlacklist_5":8,"countryBlacklist_6":9,"countryBlacklist_7":10,"countryBlacklist_8":11,"countryBlacklist_9":12,"countryBlacklist_10":13,"countryBlacklist_11":14,"countryBlacklist_12":15,"countryBlacklist_13":16,"countryBlacklist_14":17,"countryBlacklist_15":18,"countryBlacklist_16":19,"birthdayClaimIssuerId":20,"birthdayClaimIssuerBBJIdenState":21,"currentYear":22,"currentMonth":23,"currentDay":24,"minAge":25}`

	// KycBySignaturesVerificationKey is verification key to verify kycBysSignature circuit
	KycBySignaturesVerificationKey VerificationKeyJSON = `{"protocol":"groth16","curve":"bn128","nPublic":26,"vk_alfa_1":["20491192805390485299153009773594534940189261866228447918068658471970481763042","9383485363053290200918347156157836566562967994039712273449902621266178545958","1"],"vk_beta_2":[["6375614351688725206403948262868962793625744043794305715222011528459656738731","4252822878758300859123897981450591353533073413197771768651442665752259397132"],["10505242626370262277552901082094356697409835680220590971873171140371331206856","21847035105528745403288232691147584728191162732299865338377159692350059136679"],["1","0"]],"vk_gamma_2":[["10857046999023057135944570762232829481370756359578518086990519993285655852781","11559732032986387107991004021392285783925812861821192530917403151452391805634"],["8495653923123431417604973247489272438418190587263600148770280649306958101930","4082367875863433681332203403145435568316851327593401208105741076214120093531"],["1","0"]],"vk_delta_2":[["19128064920055492247805374985983083930165378749179480909913816724406615335296","19288160050156725733802379244448689099841958569866762821564854312784709506393"],["17428331044619925399916395861726985657968925788307219189262474327320267761426","2419828576521667459164020710642577668561374285085862557553780257161628012160"],["1","0"]],"vk_alphabeta_12":[[["2029413683389138792403550203267699914886160938906632433982220835551125967885","21072700047562757817161031222997517981543347628379360635925549008442030252106"],["5940354580057074848093997050200682056184807770593307860589430076672439820312","12156638873931618554171829126792193045421052652279363021382169897324752428276"],["7898200236362823042373859371574133993780991612861777490112507062703164551277","7074218545237549455313236346927434013100842096812539264420499035217050630853"]],[["7077479683546002997211712695946002074877511277312570035766170199895071832130","10093483419865920389913245021038182291233451549023025229112148274109565435465"],["4595479056700221319381530156280926371456704509942304414423590385166031118820","19831328484489333784475432780421641293929726139240675179672856274388269393268"],["11934129596455521040620786944827826205713621633706285934057045369193958244500","8037395052364110730298837004334506829870972346962140206007064471173334027475"]]],"IC":[["939914628341458097516560836826640061417412666479488218276452351605045261272","15618548428001333556505887616624363833376841092682878676319870665446762679381","1"],["19702481984208948066905921949401701892622081248277421181828805140556810527961","591249296144528568644572307700485352989887758126147542074984634824580349790","1"],["21597213372011691592118423761747260467136980646123047205544845233948785583454","5196511894272462787215205652324772860730227888212475230176185917563588975089","1"],["19947365043590033317698858174466422267665557198488278124968639335975352349566","6888103996945586203623646000310666601467283622375764292824085463872350533111","1"],["929770654496063678443820024355937299267629214892684305904187689767724951073","6221018702137222207323223468439959521912038660689132882244771195851036027436","1"],["15310074104932801501644246260897669637226023081793799067952914365167212169682","5879166969760669151358319200626715484518003603727122733880632094958992877599","1"],["17971643265209741168572406785554242786918944839698650229434685581374741153049","3868842795321072870819307332638694135310027298425990065378780298275327376227","1"],["495745971801346882197336692718840760523647268109405134318265470241548603318","21130030807975271815094195118605945442199078822769333781275642658294010578431","1"],["15761017688632999952168047278435514149120694920182536586342231270126590704103","21392966441954723879763377537982650705338045410646499550683341262986076462497","1"],["14761637176340383428252649965810248295683877890169419113056281631580002773413","6731526493847224925576121390364151881861699044372365347469588873402831839941","1"],["15851835685574472775862827117616441786754152565801802145266865670177972669062","15874065945662487959415069605527499236937033588760848794428995176829519629867","1"],["17373726396273537796870417619744542499830071779785286088433182026540191523591","9161378079365747151628520286205749938233800006194631547321283555989284857225","1"],["13312905596080705851601635638064751396651919659332733234631412807907395745857","1260373832423040695616485404682224203278851493376258227901013628705543689541","1"],["6131427624040614882789546287840649427398137387259427267973401956189842853429","3972267586492425100185156855077534857818312507309978391302972105444087606453","1"],["21198228523344347205296190384150182495468439418409269811101629880033140307176","15096741571751104933099141427282568324191338977008173304759099173470273098928","1"],["3222590932508072780326303585610580402925045277262286309448642463994919513778","16639917920181277326312669550618803226863219580774101180998300717248211138561","1"],["2087448661411056266461299935931047190539985050070525918870157796554208345315","7856861375719716246409097742827048670200813977983835009772546370489019665395","1"],["19937371855088628377463856848053459016326089130102447720252506729674313901644","8472010261302839748150687516534356087779781288686922969203004672254591955186","1"],["554270891896844532066866172748318370009985132578090355693920520674291649645","10699156451760641087986249338477737746414757627181053986536524135969186196189","1"],["5465061226051267980642510344843071304798020617742311187173887598903006967820","11893193670297230058362002712839315440501820807975395434613315133500764621711","1"],["14641687371361924497352317469887183845134509606912861194557692030600273165435","17270925897527243275922678840690882543817991761890068516463770595105142817757","1"],["17297740399105498043635452523970204725442613466771891216931543851047613222062","11970135855505351357087071704234862227389490581222449923825910728832852829134","1"],["13963696747816354569374969450817020976298346550369939668146041309950121851602","3334597054598802667360397783523455682253476471939428640736044957691532983596","1"],["16921733156938179995722747086695678352508229050490088523030962286491935657587","19580663880552112182040173817092822585729774224638798300122624436719055011580","1"],["3332736485186613668142732636026375877311430630177998823651523088195510999347","8443620575343412122110287749388130100175803277925508709171151849001928250366","1"],["11005060412361570098083562497542120757858345364067259480244409487317856148376","20901994404640877101828132182462220444734163309092115188931910219812283716947","1"],["1367787598278437985643819431598330931335252418212065257020911130676301915360","7732341683215070702836276408170949477995243447094745409006552425064825017753","1"]]}`
)

// KYCBySignatures represents KycBySignatures circuit
type KYCBySignatures struct {
}

// nolint // common approach to register default supported circuit
func init() {
	RegisterCircuit(KycBySignaturesCircuitID, &KYCBySignatures{})
}

// GetVerificationKey returns verification key
func (c *KYCBySignatures) GetVerificationKey() VerificationKeyJSON {
	return KycBySignaturesVerificationKey
}

// GetPublicSignalsSchema returns schema to parse public signals
func (c *KYCBySignatures) GetPublicSignalsSchema() PublicSchemaJSON {
	return KycBySignaturesPublicSignalsSchema
}

// PrepareInputs rerurns inputs for circuit KycBySignatures
func (c *KYCBySignatures) PrepareInputs(in TypedInputs) (map[string]interface{}, error) {

	kycInputs, ok := in.(KYCBySignaturesInputs)
	if !ok {
		return nil, errors.New("wrong type of input arguments")
	}
	ageClaimInputs, err := c.prepareRegularClaimInputs(
		kycInputs.KYCAgeCredential, kycInputs.KYCAgeCredentialRevocationStatus,
		"birthday", kycInputs.AgeSignatureProof)
	if err != nil {
		return nil, err
	}
	countryClaimInputs, err := c.prepareRegularClaimInputs(
		kycInputs.KYCCountryOfResidenceCredential,
		kycInputs.KYCCountryOfResidenceRevocationStatus, "country",
		kycInputs.CountrySignatureProof)
	if err != nil {
		return nil, err
	}

	authClaimInputs, err := c.prepareAuthClaimInputs(kycInputs.ID,
		kycInputs.IssuerAuthClaimMTP, kycInputs.IssuerAuthClaimClamTreeRoot,
		kycInputs.Challenge, kycInputs.Signature, kycInputs.PubKey)
	if err != nil {
		return nil, err
	}

	publicInputs, err := c.prepareCircuitPublicInputs(kycInputs.Rules)
	if err != nil {
		return nil, err
	}
	inputs := mergeMaps(ageClaimInputs, countryClaimInputs, authClaimInputs, publicInputs)
	return inputs, nil
}

type SignatureProof interface {
	signatureProofMarker()
}

type BaseSignatureProof struct {
	IssuerID        *core.ID
	IssuerTreeState TreeState
	Siblings        []*merkletree.Hash
}

type BJJSignatureProof struct {
	BaseSignatureProof
	IssuerPublicKey *babyjub.PublicKey
	Signature       *babyjub.Signature
}

func (BJJSignatureProof) signatureProofMarker() {}

// prepareRegularClaimInputs prepares inputs for regular claims
func (c *KYCBySignatures) prepareRegularClaimInputs(claim Claim,
	rs RevocationStatus, fieldName string,
	signatureProof2 SignatureProof) (map[string]interface{}, error) {

	inputs := make(map[string]interface{})
	var err error

	inputs[fieldName+"Claim"] = bigIntArrayToStringArray(claim.ZKInputs)

	switch sp := signatureProof2.(type) {
	case BJJSignatureProof:
		inputs[fieldName+"ClaimIssuerBBJClaimMtp"] = bigIntArrayToStringArray(
			PrepareSiblings(sp.Siblings, LevelsKYCCircuits))
		inputs[fieldName+"ClaimIssuerBBJAx"] = sp.IssuerPublicKey.X.String()
		inputs[fieldName+"ClaimIssuerBBJAy"] = sp.IssuerPublicKey.Y.String()
		inputs[fieldName+"ClaimSignatureR8x"] = sp.Signature.R8.X.String()
		inputs[fieldName+"ClaimSignatureR8y"] = sp.Signature.R8.Y.String()
		inputs[fieldName+"ClaimSignatureS"] = sp.Signature.S.String()
		// Issuer identifier
		inputs[fieldName+"ClaimIssuerId"] = sp.IssuerID.BigInt().String()
		inputs[fieldName+"ClaimIssuerBBJClaimClaimsTreeRoot"] = sp.
			IssuerTreeState.ClaimsRootStr()
		inputs[fieldName+"ClaimIssuerBBJClaimRevTreeRoot"] = sp.
			IssuerTreeState.RevocationRootStr()
		inputs[fieldName+"ClaimIssuerBBJClaimRootsTreeRoot"] = sp.
			IssuerTreeState.RootOfRootsRootStr()
		inputs[fieldName+"ClaimIssuerBBJIdenState"] = sp.
			IssuerTreeState.StateStr()
	default:
		return nil, errors.New("signature type is not supported")
	}

	err = handleRevocationStateInputs(rs, fieldName, inputs)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// prepareAuthClaimInputs prepare inputs for authorization (ID ownership)
func (c *KYCBySignatures) prepareAuthClaimInputs(id *core.ID,
	mtp Proof, claimTreeRoot *merkletree.Hash, challenge int64,
	sig *babyjub.Signature,
	pubKey *babyjub.PublicKey) (map[string]interface{}, error) {

	if sig == nil {
		return nil, errors.New("signature is null")
	}

	if pubKey == nil {
		return nil, errors.New("public key is null")
	}

	inputs := make(map[string]interface{})
	inputs["id"] = id.BigInt().String()
	inputs["challenge"] = strconv.FormatInt(challenge, 10)
	inputs["BBJClaimMtp"] = bigIntArrayToStringArray(
		PrepareSiblings(mtp.Siblings, IDStateLevels))
	inputs["BBJClaimClaimsTreeRoot"] = claimTreeRoot.BigInt().String()
	inputs["BBJAx"] = pubKey.X.String()
	inputs["BBJAy"] = pubKey.Y.String()
	inputs["challengeSignatureR8x"] = sig.R8.X.String()
	inputs["challengeSignatureR8y"] = sig.R8.Y.String()
	inputs["challengeSignatureS"] = sig.S.String()

	inputs["BBJClaimRevTreeRoot"] = merkletree.HashZero
	inputs["BBJClaimRootsTreeRoot"] = merkletree.HashZero

	return inputs, nil
}

// prepareCircuitPublicInputs prepares input for public rules
// nolint:dupl // allows to change public inputs for circuit later
func (c *KYCBySignatures) prepareCircuitPublicInputs(rules map[string]interface{}) (map[string]interface{}, error) {

	inputs := make(map[string]interface{})

	countryBlackList, ok := rules["countryBlacklist"]
	if !ok {
		return nil, errors.New("country list is not provided in rules argument")
	}
	countryCodes := make([]*big.Int, 16)
	for i := range countryCodes {
		countryCodes[i] = new(big.Int).SetInt64(0)
	}
	for i, code := range countryBlackList.([]interface{}) {
		countryCodes[i] = new(big.Int).SetInt64(int64(code.(float64)))
	}
	inputs["countryBlacklist"] = bigIntArrayToStringArray(countryCodes)

	currentYear, ok := rules["currentYear"].(float64)
	if !ok {
		return nil, errors.New("currentYear is not provided in rules argument")
	}
	currentMonth, ok := rules["currentMonth"].(float64)
	if !ok {
		return nil, errors.New("currentMonth is not provided in rules argument")
	}
	currentDay, ok := rules["currentDay"].(float64)
	if !ok {
		return nil, errors.New("currentDay is not provided in rules argument")
	}

	minAge, ok := rules["minAge"].(float64)
	if !ok {
		return nil, errors.New("minAge is not provided in rules argument")
	}

	inputs["currentYear"] = new(big.Int).SetInt64(int64(currentYear)).String()
	inputs["currentMonth"] = new(big.Int).SetInt64(int64(currentMonth)).String()
	inputs["currentDay"] = new(big.Int).SetInt64(int64(currentDay)).String()
	inputs["minAge"] = new(big.Int).SetInt64(int64(minAge)).String()

	return inputs, nil
}

// KYCBySignaturesInputs represents input data for kyc and kycBySignatures circuits
type KYCBySignaturesInputs struct {
	KYCAgeCredential                      Claim
	KYCAgeCredentialRevocationStatus      RevocationStatus
	AgeSignatureProof                     SignatureProof
	KYCCountryOfResidenceCredential       Claim
	KYCCountryOfResidenceRevocationStatus RevocationStatus
	CountrySignatureProof                 SignatureProof
	ID                                    *core.ID
	IssuerAuthClaimMTP                    Proof
	IssuerAuthClaimClamTreeRoot           *merkletree.Hash
	Challenge                             int64
	PubKey                                *babyjub.PublicKey
	Signature                             *babyjub.Signature

	Rules map[string]interface{}

	TypedInputs
}