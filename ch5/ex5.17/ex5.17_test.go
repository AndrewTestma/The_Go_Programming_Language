package ex5_17

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestElementByTagName(t *testing.T) {
	input := `
<html>
<body>
	<p class="something" id = "short">
		<span class="special">hi</span>
	</p>
	<br/>
div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/04/14/What-is-a-Photographer'>&#x201c;Photographer&#x201d;?</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>Every&#xad;one takes pic&#xad;tures ev&#xad;ery&#xad;where now, 24/7/365. So does &#x201c;photographer&#x201d;, in the am&#xad;a&#xad;teur sense, still mean any&#xad;thing? I have pic&#xad;tures and ques&#xad;tions that say it does</span><span class='el'><a href='/ongoing/When/201x/2016/04/14/What-is-a-Photographer'>&#xa0;...</a></span><br/> <i>[6 comments]</i>&#xa0;&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/04/11/Audio-repair'>Speaker Dust Cap Dent Repair</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>As pre&#xad;vi&#xad;ous&#xad;ly not&#xad;ed in this space, I&#x2019;m a de&#xad;ranged au&#xad;dio&#xad;phile, and for some  years my speak&#xad;ers of choice have been from <a href='https://totemacoustic.com/en/'>Totem</a>, out of Montr&#xe9;al. In a  re&#xad;cent ren&#xad;o&#xad;va&#xad;tion a woofer got a dent in a <a href='https://en.wikipedia.org/wiki/Dust_cap'>dust cap</a>, where by &#x201c;dent&#x201d; I mean it was pushed in. I&#x2019;m post&#xad;ing the so&#xad;lu&#xad;tion here in the hopes that fu&#xad;ture searchers will find it</span><span class='el'><a href='/ongoing/When/201x/2016/04/11/Audio-repair'>&#xa0;...</a></span><br/> <i>[4 comments]</i>&#xa0;&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/04/07/Graph-it'>Getting the Picture</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>It&#x2019;s like this: Aver&#xad;ages are your en&#xad;e&#xad;my be&#xad;cause they hide change. Mak&#xad;ing graphs is easy and cheap and some&#xad;times they un&#xad;cov&#xad;er se&#xad;cret&#xad;s; more of us should do it more</span><span class='el'><a href='/ongoing/When/201x/2016/04/07/Graph-it'>&#xa0;...</a></span><br/> <i>[1 comment]</i>&#xa0;&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/04/01/Cool-Aussie-Phoners'>Cool Aussie Phoners</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>What the ti&#xad;tle says: Three pho&#xad;tos of Aussies hold&#xad;ing phones look&#xad;ing cool</span><span class='el'><a href='/ongoing/When/201x/2016/04/01/Cool-Aussie-Phoners'>&#xa0;...</a></span><br/>&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/03/31/Serverlessness'>Serverlessness</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>To&#xad;day Mi&#xad;crosoft an&#xad;nounced <a href='https://azure.microsoft.com/en-us/services/functions/'>Azure Func&#xad;tions</a>, join&#xad;ing <a href='https://cloud.google.com/functions/docs'>Google Cloud Func&#xad;tions</a>    and (from 2014) <a href='https://aws.amazon.com/lambda/details/'>AWS Lamb&#xad;da</a>. This is fun stuff, and might be a big deal</span><span class='el'><a href='/ongoing/When/201x/2016/03/31/Serverlessness'>&#xa0;...</a></span><br/> <i>[3 comments]</i>&#xa0;&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/03/28/Mobile-Aerials'>Mobile Aerials</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>A lit&#xad;tle while ago I tweet&#xad;ed &#x201c;One thing phone-cams aren&#x2019;t much good for is shoot&#xad;ing out air&#xad;plane windows.&#x201d; Since then, I&#x2019;ve no&#xad;ticed my Nexus 5X look&#xad;ing at me in a hurt sort of way</span><span class='el'><a href='/ongoing/When/201x/2016/03/28/Mobile-Aerials'>&#xa0;...</a></span><br/> <i>[1 comment]</i>&#xa0;&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/03/27/Practical-photography'>Lenses and Cameras in 2016</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>I&#x2019;m on the way back from a cou&#xad;ple of weeks in Aus&#xad;trali&#xad;a, and of course Pic&#xad;tures Were Tak&#xad;en. I brought al&#xad;most all my photo-gear but used it very un&#xad;even&#xad;ly; con&#xad;clud&#xad;ed that I have too many lens&#xad;es, and was left won&#xad;der&#xad;ing whether you re&#xad;al&#xad;ly even need a cam&#xad;era any more. Here&#xad;with notes il&#xad;lus&#xad;trat&#xad;ed with Pacific-ocean (most&#xad;ly) pic&#xad;tures</span><span class='el'><a href='/ongoing/When/201x/2016/03/27/Practical-photography'>&#xa0;...</a></span><br/> <i>[4 comments]</i>&#xa0;&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/03/18/Parrot-story'>Galah Story</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>We&#x2019;re vis&#xad;it&#xad;ing friends in Aus&#xad;tralia and I watched a pair of <s>par&#xad;rots</s>   in&#xad;ter&#xad;act  (oop&#xad;s, Mar&#xad;ius Coomans writes from Aus&#xad;tralia to tell me they&#x2019;re Galah&#xad;s);  pho&#xad;tographed them, but didn&#x2019;t un&#xad;der&#xad;stand. Oh, and a koala</span><span class='el'><a href='/ongoing/When/201x/2016/03/18/Parrot-story'>&#xa0;...</a></span><br/> <i>[2 comments]</i>&#xa0;&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/03/07/City-Thinking'>Urbanity</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>Ci&#xad;ties are our rule now, any&#xad;thing else the ex&#xad;cep&#xad;tion.  I&#x2019;m bik&#xad;ing most work&#xad;days, on con&#xad;crete over the ocean in&#xad;to the  stone heart of a small big city, get&#xad;ting ten dozen chan&#xad;nels of non&#xad;stop  ur&#xad;ban in&#xad;put and ev&#xad;ery day I won&#xad;der where we&#x2019;re all go&#xad;ing.  The fu&#xad;ture is  dis&#xad;tribut&#xad;ed un&#xad;even&#xad;ly and&#x2009;cities con&#xad;cen&#xad;trate the  un&#xad;even&#xad;ness</span><span class='el'><a href='/ongoing/When/201x/2016/03/07/City-Thinking'>&#xa0;...</a></span><br/> <i>[1 comment]</i>&#xa0;&#xa0;</div>
<div class='os'><span class='ot'><a href='/ongoing/When/201x/2016/02/28/Lightroom-Mobile-Nexus'>Lightroom, Mobile, Nexus</a></span>&nbsp;&#xb7;&nbsp;<span class='ofp'>In which I re&#xad;port on us&#xad;ing the Nexus 5X in RAW mod&#xad;e, with the help of  Adobe Light&#xad;room, and on work&#xad;flows for mo&#xad;bile pho&#xad;togs.  With il&#xad;lus&#xad;tra&#xad;tions from Vancouver&#x2019;s <a href='http://westvancouver.ca/parks-recreation/parks/lighthouse-park'>Light&#xad;house Park</a></span><span class='el'><a href='/ongoing/When/201x/2016/02/28/Lightroom-Mobile-Nexus'>&#xa0;...</a></span><br/> <i>[3 comments]</i>&#xa0;&#xa0;</div>
<body>
</html>
`
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}
	for _, data := range ElementByTagName(doc, "span") {
		fmt.Printf("%+v\n", data.Data)
	}
}
