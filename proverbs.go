package proverbs

import (
	"fmt"
)

const a = `Tb Cebireof, ol Ebo Cvxr

Qba'g pbzzhavpngr ol funevat zrzbel, funer zrzbel ol pbzzhavpngvat.
Pbapheerapl vf abg cnenyyryvfz.
Punaaryf bepurfgengr; zhgrkrf frevnyvmr.
Gur ovttre gur vagresnpr, gur jrnxre gur nofgenpgvba.
Znxr gur mreb inyhr hfrshy.
vagresnpr{} fnlf abguvat.
Tbszg'f fglyr vf ab bar'f snibevgr, lrg tbszg vf rirelbar'f snibevgr.
N yvggyr pbclvat vf orggre guna n yvggyr qrcraqrapl.
Flfpnyy zhfg nyjnlf or thneqrq jvgu ohvyq gntf.
Ptb zhfg nyjnlf or thneqrq jvgu ohvyq gntf.
Ptb vf abg Tb.
Jvgu gur hafnsr cnpxntr gurer ner ab thnenagrrf.
Pyrne vf orggre guna pyrire.
Ersyrpgvba vf arire pyrne.
Reebef ner inyhrf.
Qba'g whfg purpx reebef, unaqyr gurz tenprshyyl.
Qrfvta gur nepuvgrpgher, anzr gur pbzcbaragf, qbphzrag gur qrgnvyf.
Qbphzragngvba vf sbe hfref.
Qba'g cnavp.`


func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func init(){
	p := make([]byte, len(a))
	for i:= 0; i < len(a); i++ {
		p[i] = rot13(a[i])
	}
	fmt.Println(string(p))
}