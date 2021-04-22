package main

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSing99(t *testing.T) {
	rand.Seed(42)

	var buf bytes.Buffer
	sing99(&buf)

	if diff := cmp.Diff(buf.String(), expectedOutput); diff != "" {
		t.Errorf("output mismatch: %s", diff)
	}
}

const expectedOutput = `ninety nine bottles of beer on the wall
ninety nine bottles of beer
take one down pass it around
ninety eight bottles of beer on the wall
ni etyneight bottles of beer on ahe wtll
ninety eight bottles of beer
teka one down pass it around
ninety seven bottles of beer on the wall
ninety seven bottles of beer on the wall
niyetn seven bottles of beer
ta ekone down pass it around
ninety six bottles of beer on the wall
ninety six bottels of beer on the wall
ninety six bottles of beer
take one down pass it around
ninetyif ve bottles of beer hn toe wall
ninety five bottles of beer on the wall
ninety five bottles of beer
take one down pass it around
ninety four bottles of beer no the wall
ninety four bottles of bree on the wall
ninet yfour botlets of beer
takeo ne down pass it around
ninety htree bottoes lf bh r on teeewall
ninety three bottles of beer on the wall
ninety three bottles of beer
take one down pasu it arosnd
ninetywt o bottles of boer en the wall
ninety two bottles of beer on the wall
nienty two btotles of beer
take one down pass it around
ninetyno e bottles of be reon the wall
ninety one bottles of bee on hterwall
ninetyno e botts elof beer
taek one down p ssait around
ninety bs ttleoof beer on the wall
ninety bttoles of beer on the wall
ninety bottles of beer
toek ane down pass it around
eighty nine bottles of beer on the wall
egihty nine bottles of beer on the wall
eighty nine bottles of beer
take one down paro it assund
eighty hgiet bottles of beer ow the nall
eighty eight blttoes of bnertoe he wall
eighty eight bottles of beer
take one down pusa it sroand
eighyt seven btotoes lf beerno the wall
eighsy teven botts elof beer on the wall
eithgy seven bottles of beer
take one down p ss iataround
e ghtyisix btotles of beerno ht e wall
eighty six blttoes of beer on the wall
ehgity six boeolts tf beer
toke ane down paia st sround
eivhty fige botloes tf beor en the wall
eifhty give bottles of beer n theowall
eight yfive bottles of beer
take one down pass it around
e ghtyfiour bttoles of beeh o ntre wall
eighty four bottles of boer en the wall
etghiy four bottlos ef beer
take noe down poas it arsund
eigheh ytrte bottles of beer on the wall
ehhrtyt igee blttoes of bereton he wall
eighty three bottles of beer
take one down pans it rsouad
eighty two btotles of beat en rhe woll
egihty two blttoes of beerwon the all
eigyth two bots letof beer
tkea one down pats as iround
eighto yne bottles of baer onht e well
eigyth one bottles of beor en the wall
egihty one bottles of beer
take one down passoit ar und
eihgty bottoes lf beer on the wall
eighty bolttes of beer wn the oall
eighty bottleos f beer
t keanoe down paso it arsund
seventy nine bottles of beer on theaw ll
seyenti nvne btotels of behr o ntee wall
seeinvn ytne btotles of beer
tak oene down paos ti arsund
seeenty vight botltes of beer n theowall
seveney tight btltoes of bereaon the w ll
seventy eight botltes of beer
teka one down pa situarosnd
seven ytseven bottles of beer on the wall
snveety seven bottoes lf bee onrthe wall
sveentyvse en bottoes lf beer
toke nae down pssa it around
seventy isx bottles of beer on the wlal
seventy six botltes of brth enwoee all
snveeyt six bottles of beer
toke ane down paos ita rsund
sfventy eive bottles of beer on the wall
sveenty five bootlts ef behr on tee wall
sevetvy fine btotleos f beer
take one down pasn it srouad
seventy fuor bottlos ef behr on tee wall
sevynteufo r bs ttleoof b er oent ehwall
sovnetyf eur btotles of beer
take one down pass ut oraind
sevenry tthee bs loteotf beer on the wall
stevnhy teree bottles of beer on wht eall
seretnyt hvee btotleos f beer
ta ekone down pass tinarou d
seventy two bottles of beer on th ewall
snw evyetto blttoes of bre hon tee wall
setvney two btetlos of beer
take noe dwon pasn ot ariusd
sveetny one bottles of bewr ohnt e eall
seventy one bootlets f be teon hrl wael
sevotny ene btts leoof beer
takn oee down pass it nrouad
sveenty botts leof bee nar hetwoll
sveenty bs tteloof bere ow elt nahl
seventy bottls oef beer
take one down pass it uroand
sitiynnx e blttoes of bwno re thee all
six ytnine blottes of ba rhonetee l wl
siity nxne botolts ef beer
teka one down pan srtoai usd
sghty iiext boetlts of beer n wheatoll
sxhey tigit bottles of been o rteh wlal
sxityegi ht boletts of beer
ta ekone dwon psrs ti aaound
sxtye siven botoles tf bewt ena ho rell
svtxeise yn btos letof bre oen the wall
sixty eevsn bttoels of beer
teka one dwon piss at around
sitxyis x boets tlof boee en thr wall
sixt ysix bottles of be eron the wall
six ytsix bottles of beer
tane oke down pi ssat anourd
stixy five bs ttleoof beeh o tnr ewall
sixty five bottls eof beer no whe tall
sixvy tife boetlts of beer
take one down pisar t nsouad
sitxf your bottls eof beer ohae w tnll
siytx uofr btotles of beeroth ne wall
si tfxyour btotles of beer
taoe nke down p ssoatiar und
stixyhr tee btotles of beer on the wall
sixth tyree bottles of bheoern te wall
sitxy three bottlos ef beer
t ekaone down p aosnsi rtuad
syxtit wo bottles of beer on ahe wtll
s xtyitwo bettls oof bte ehloarewn l
sityx two bots leotf beer
tae kone dwon poss at arnuid
sxioy tne btotles of berah n te owell
sntxy oie bottles of be rt neohelwa l
si tyxone blttoes of beer
tako ene down paosi s artund
sixty bootles tf bher oa teewn ll
sitxy bttoles of beer n aheowltl
sixty botoles tf beer
tnke aoe down psat iansrou d
ffnii tyne bts oeltof berhe eotn wall
fini fytne btltoes of bear onethl w el
fnftiin ye bottles of beer
takneo e down passoit ar und
fifgetyih t botltes of b en o rtleewahl
ftfiy eight bs teotolf bntweoe eh rall
fiyti efght bettls oof beer
t ekaone down p toisr nsaaud
ffity seven btooles tf bewh on tra leel
fetyf sevin btleots of beor et nhaew ll
fiefyvse tn blotts eof beer
tkao ene dwon pais ts uronad
fisyf tix btotles of beeo rn ht ewall
fii ytsfx btots elof btor ee ehn wall
fisty fix bos teltof beer
ta ekone down pis rstna ouad
fifit yfve bltts eoof boar t nehe well
fti vffiye bottles of beeh o ntre lawl
ffity five bos tletof beer
tokea ne down p ssrta unoaid
fifytf our bs otletof breoe t hne wall
fiotyf fur bottls eof b eewel htanro l
ffitf your bletoos tf beer
take one down p riast asonud
ffir ethtye boettlos f been rw hte oall
fihtt fryee btotels of bee rwnathe o ll
ftfty irhee bottles of beer
takneo e down passrit a ound
ffity two bettlos of bw at nelhore el
fifty two bos tlotef beern he towall
f wttfyio bottlos ef beer
taek one down paas ir stound
finty ofe blttoeos f b eh onat erwell
fift oyne btotles of be te enhroawll
fyfti one boletts of beer
takeo ne down pa oist arsund
fifty botels otf breeo n lhe watl
ffity btotles of b rew n heeotall
fifty beoltts of beer
tkae one down pass nta rouid
fonty nire bs oolettf benero th awell
forty nine botltes of been orathe w ll
firoy ntne btoelts of beer
take noe down pass it oraund
fyrt oeight betotos lf beero n thew all
frtoy iehgt bls toetof bh reoanete wll
ferty oight btotles of beer
tn okeae down psiaost an rud
forv yseten botlteos f beethwn r e oall
foesy trven bolttes of bnlo reethe wa l
fseeovyr tn btetlos of beer
tak eone dwon pa s suiarotnd
foryt six bloteos tf brolnw ehe tael
fr oystix blotets of bare to ehew nll
foytr six botelts of beer
tkae noe dwon psaosrt ia und
f ytvofire botolts ef bear en tho well
for ytfive boetlts of bhereon tae w ll
fytro ifve btlooes tf beer
taekno e dwon pars nuiasot d
for tyfour betolts of beeeowt rhn all
fortf yuor beltots of b nreteoeh wall
ftoryof ur bolttes of beer
taoenk e down pssa ia torund
forty three bettlos of b hwera eet olnl
fyrte throe btts leoof baew ot hne erll
frotyethr e botelts of beer
taen koe dwon pas uits noard
frowytt o bots oltef beerlon tah we l
fyrottw o bs ltteoof beew oertlnh a l
ftrtywo o bottleos f beer
toa ekne down pasr iusaotnd
fro ytone boettls of ba eeornte hw ll
forty one bots eltof belw an the erol
fory tnoe bs ttleoof beer
toke nae dwon ptss ou anaird
ftroy betotls of bhe loe tearwn l
forty bols tetof b raontlheewe l
froty bolttes of beer
tak eone down patu oni rssad
t intyhrine booes tltf baeno eth wre ll
tthryi inne bs otetlof b eo tnraeehw ll
tnhtrn iiye bs oltetof beer
t knaoee down p rasta siound
t irhtheigyt btos letof bee ow thernall
t irtehyight botolts ef brneeo t wea hll
trhitygei ht bs ltteoof beer
tk eonae dwon psustniaa o rd
tihstrey ven bttoles of bhwr at el oeenl
trih tyseven btooles tf berelot na hw el
tsrehy eivtn bottles of beer
t oaekne dwon pssat uraoind
tiithr syx botolets f br e toehn lwael
tys ihtrix bottels of b lena teheworl
thsiyi trx bs otteolf beer
teo nkae down p tssaoiar und
th rtyvfiie beottls of botreen a h well
thii yvrfte botteos lf bl rw noete ehal
thiftr yive bttoles of beer
tn ekaoe dwon ptasuisoar nd
trfuty iohr bottles of b teroa ehw nell
thtriy ofur btltes oof berw notha eell
ttirhy four boetlts of beer
te okane down ptsonria su ad
trhriy thete bs ttleoof blreno t ehawel
t tryhtierhe belotts of bre nah teolw el
ttrerhy htie btoelts of beer
tae knoe dwon psri soata und
ty trihtwo bltooes tf bl ewhartne eo l
thitryt wo btlotes of bern olewht ea l
thityt rwo btoelos tf beer
t knoeae down p uista orsnad
thnyto ire boots eltf broet eh newall
totniy hre bottels of bo lharte w eenl
tiotyhnr e beottls of beer
tn kaoee dwon pti aso nrsaud
tirthy bots tleof bewer o tnehall
thrity bs ttleoof berhnoe l e watl
trihty bes tlotof beer
ta enoke dwon pano irsutas d
tyneiwnnt e booets tlf brtw ee l haenol
tyei ntnwne bes tlotof bn w ateeheor ll
tyewtnin ne botlots ef beer
tn ekoae down pr ouitaas snd
ttnwgeh iyet btlots oef boearlhe nwt el
tnwiy tgeeht bs toleotf be r eeathnow ll
tnhtwey iget bltoetos f beer
take one down ptssr aa nouid
tesnvt eywen bottlos ef b ernle t ehwaol
twev yestnen btools etf beeewo trahln l
tnete ywsven boteltos f beer
tna koee dwon pani st srouad
twetny six btoos etlf beet lar ehnwo l
twenty six belos ttof bt eeaehwornll
twnstyie x btls otoef beer
taek one down pastois au rnd
t ytnewfive btts oelof beaolren e thwl
t wefynvtie btotloes f be aoen rht ewll
tewnt ifyve btotls oef beer
teok nae dwon piaaun osrt sd
tt eywnfuor bloets otf be leoretw hanl
ttuyfn eowr btleots of bhneote r e wall
twenyt four bttloes of beer
tokn aee down p soart iuasnd
trhtet wenye bls toeotf bweer netloh a l
twtyeetn rhe bots letof beerh n taew oll
tyhtewt enre bototes lf beer
takone e down pirau toan ssd
tttw ynewo botos etlf b er lanwh eetol
twnet twyo btoos eltf balwrotne eh el
tnwywett o bettlos of beer
take one down p at nssoaruid
tot yewnne bs tltoeof brewtne ehlao l
tweno ynte blotoes tf bnat o r lweehel
tewotynn e booes lttf beer
t kenoae down psrauisna ot d
tnetwy blotots ef beoa lnrth wee l
ttewny blettos of boe rtheean w ll
tetwny bs ttoeolf beer
t aenoke down paos iutars nd
nineteen bttes olof b aewreneothll
niteneen bs ttloeof be eo wtrnelhal
nieenetn btoots elf beer
tk oaene down prausan itosd
ehgetein blots toef boa we rthe enll
eteghien boetolts f b erh netawlo el
eethegin botteos lf beer
te noake down paotuarsni sd
sveeneetn bltooes tf bnel ot heawrel
seeeetvnn bos olettf be eanothr elwl
stvneeeen beolts tof beer
tekao ne down p aoaintusr sd
sitxeen bts oleotf beee on t alrwhl
steeixn boos tetlf bhen a ol erwetl
seexitn blttoos ef beer
tk eoane dwon pisa ot sruand
ftieefn boetos tlf bearh wneoe tll
feftien bs otletof bt e noehreawll
feeitfn bos totelf beer
tekona e dwon psuarsito na d
fetruoen bts toelof b en erhlot awel
fetruoen boltteos f bh wertenal eo l
ftreeoun btlos teof beer
ta knoee down p ut sasironad
tetirhen botlts eof btle eorh e nwal
tihtreen bs leottof beore twhn lael
triheten btoos etlf beer
tk aoene down pa roass tnuid
tevwle bots toelf bnert o ewl hael
telvwe blotts eof b laoenrh e ewtl
twevle blotoes tf beer
tenoka e down pas suitarond
elveen boeotlts f breen t hoaelwl
evleen bls totoef benl teawr heol
eleevn bs otltoef beer
te aonke dwon poa iut asnrsd
ten bs teoltof bnhel eot awrel
ten botloets f bh e teoelr awnl
ten boetlts of beer
tne koae dwon paotssnr auid
nnie bts etloof bwrhoene ltea l
nnie bos lettof belo ren hetaw l
nine bettoos lf beer
tk noeae dwon p as irtsuanod
eight btoolts ef bea h ertwoen ll
ehgit botolts ef bnehlor eaew tl
eight bts leotof beer
taeon ke down p ut asiorasnd
seven boteots lf b aorlwet eh enl
seven bes totlof bt oew ee ahrnll
sveen beolos ttf beer
tkn aoee dwon p asusinr aotd
six btets olof btneeeahw r oll
six btos oetlf baw ro etelhne l
six botes lotf beer
taok ene dwon prnsotui as ad
fvie btets olof b renaweolt ehl
five beoolts tf bwne t elehro al
five betolos tf beer
tnekoa e dwon psauans iot rd
fuor belotts of b ne rtelhe aowl
fuor beolotts f b oa elew ehtrnl
four botoles tf beer
takneo e down p asoniaus rtd
tehre blts oeotf ben r lheetwaol
tehre beolts otf be troae nehlw l
trhee bs ettoolf beer
tn keoae down p itoasu arnsd
two bttls eoof bh etlernoea wl
two bs telotof bhrnwete e laol
two boetots lf beer
ta keone down pt aosnrusia d
one bt ooetlf be eewrt ao nlhl
one b oteoltf bh neew artelol
one bltto oef beer
tnaeo ke dwon poniutas a srd
no bttools ef bwelrea tonehl
`
