package handlers

import (
	"fmt"
	sl "github.com/djosephsen/slacker/slackerlib"
	"math/rand"
	"time"
	"net/http"
	"encoding/json"
	"regexp"
)

type insult struct{
   Insult   string
   Severity string
}

var LoveAndWar = sl.MessageHandler{
	Name: `LoveAndWar`,
	Usage: `<botname> (love|insult) <noun>: bot replies with a compliment or insult respectively ** Warning this plugin uses external API's that may return NSFW responses**`,
	Method:  `RESPOND`,
	Pattern: `(?i)(love|insult) (@*\w+)`,
	Run: func(e *sl.Event, match []string){
		var reply string
		act:=match[1]
		user:=match[2]
		if isme,_ := regexp.MatchString(`(?i)me`,user); isme{
			user = e.Sbot.Meta.GetUserName(e.User)
		}
		now:=time.Now()
		rand.Seed(int64(now.Unix()))
		if isLove,_ := regexp.MatchString(`(?i)love`,act); isLove{
			reply=makeLove(user)
		}else if isWar,_ := regexp.MatchString(`(?i)insult`,act); isWar{
			reply=makeWar(user)
		}
		e.Respond(reply)
	},
}

func makeWar(user string) string{
	it1:=[]string{
	`lazy`,
	`stupid`,
	`californian`,
	`slimy`,
	`smelly`,
	`slutty`,
	`pompous`,
	`communist`,
	`wangnose`,
	`pie-eating`,
	`racist`,
	`eliteist`,
	`fascist`,
	`drug-snarfing`,
	`slovenly`,
	`tone-deaf`,
	`ugly`,
	`buck-toothed`,
	`creepy`,
	`goat-faced`,
	}

	it3:=[]string{
	`spaz`,
	`douche`,
	`turd`,
	`ass`,
	`rectum`,
	`butt`,
	`poop`,
	`armpit`,
	`crotch`,
	`bitch`,
	`slime`,
	`prick`,
	`slut`,
	`taint`,
	`roach`,
	`snot`,
	`boner`,
	`shart`,
	`nut`,
	`sphincter`,
	}

	it2:=[]string{
	`pilot`,
	`canoe`,
	`captain`,
	`pirate`,
	`hammer`,
	`knob`,
	`box`,
	`jockey`,
	`nazi`,
	`waffle`,
	`goblin`,
	`nazi`,
	`biscuit`,
	`clown`,
	`socket`,
	`monster`,
	`clown`,
	`hound`,
	`recepticle`,
	`balloon`,
	}

	n:=rand.Intn(2)+1
	switch n {
		case 1:
	   	i:=new(insult)
			resp,_:=http.Get(`http://pleaseinsult.me/api`)
			dec := json.NewDecoder(resp.Body)
			dec.Decode(i)
			return fmt.Sprintf("Hey %s... %s",user, i.Insult)
		case 2:
			return fmt.Sprintf("%s is a %s %s %s",user, it1[rand.Intn(len(it1))], it2[rand.Intn(len(it2))],it3[rand.Intn(len(it3))])
	}
	return fmt.Sprintf("... derp, excuse me I have a bug: %s",n)
}

func makeLove(user string) string{
	love:=[]string{
		`You deserve a promotion.`,
		`I appreciate all of your opinions.`,
		`I like your style.`,
		`Your T-shirt smells fresh.`,
		`You are like a spring flower; beautiful and vivacious.`,
		`I am utterly disarmed by your wit.`,
		`I really enjoy the way you pronounce the word 'ruby'.`,
		`I like those shoes more than mine.`,
		`Nice motor control!`,
		`You have a good taste in websites.`,
		`Your mouse told me that you have very soft hands.`,
		`You are full of youth.`,
		`I like your jacket.`,
		`You have a good web-surfing stance.`,
		`You should be a poster child for poster children.`,
		`I appreciate you more than Santa appreciates chimney grease.`,
		`I wish I was your mirror.`,
		`I find you to be a fountain of inspiration.`,
		`You have perfect bone structure.`,
		`I disagree with anyone who disagrees with you.`,
		`Have you been working out?`,
		`With your creative wit, I'm sure you could come up with better compliments than me.`,
		`I like your socks.`,
		`You are so charming.`,
		`You're tremendous!`,
		`Your smile is breath taking.`,
		`How do you get your hair to look that great?`,
		`Take a break; you've earned it.`,
		`Your life is so interesting!`,
		`The sound of your voice sends tingles of joy down my back.`,
		`I enjoy spending time with you.`,
		`I would share my dessert with you.`,
		`I would love to visit you, but I live on the cloudbutts.`,
		`I love the way you click.`,
		`You're invited to my birthday party.`,
		`All of your ideas are brilliant!`,
		`If I freeze, it's not a computer virus.  I was just stunned by your intellect.`,
		`You're spontaneous, and I love it!`,
		`You should try out for everything.`,
		`You make my data circuits skip a beat.`,
		`You are the gravy to my mashed potatoes.`,
		`You get an A+ in the rollbook of my heart!`,
		`I'm jealous of the other websites you visit, because I enjoy seeing you so much!`,
		`I would enjoy a roadtrip with you.`,
		`If I had to choose between you or 17lbs of chocolate, I would choose you`,
		`I like you more than the smell of Grandma's home-made apple pies.`,
		`You would look good in glasses OR contacts.`,
		`Let's do this again sometime.`,
		`You could go longer without a shower than most people.`,
		`I feel the need to impress you.`,
		`I would trust you to pick out a pet fish for me.`,
		`I'm glad we met.`,
		`Will you sign my yearbook?`,
		`You're so smart!`,
		`We should start a band.`,
		`You're cooler than ice-skating Fonzi.`,
		`I heard you make really good French Toast.`,
		`I like your pants.`,
		`You're pretty groovy, dude.`,
		`When I grow up, I want to be just like you.`,
		`I tweeted all my friends about how cool you are.`,
		`You're so awesome, you can play any prank, and get away with it.`,
		`I can tell that we are gonna be friends.`,
		`I just want to gobble you up!`,
		`You're awesome. Treat yourself to another compliment!`,
		`You're pretty high on my list of people with whom I would want to be stranded on an island.`,
		`You could probably lead a rebellion.`,
		`:heart: :heart: :heart:`,
		`You are more fun than a Japanese steakhouse.`,
		`Your voice is more soothing than Morgan Freeman's.`,
		`You could be drinking whole milk if you wanted to.`,
		`I support all of your decisions.`,
		`You are as fun as a hot tub full of chocolate pudding.`,
		`Being awesome is hard, but you'll manage.`,
		`Your skin is radiant.`,
		`You could survive a zombie apocalypse.`,
		`I wish I could move your furniture.`,
		`You're so rad.`,
		`Your glass is the fullest.`,
		`I find you very relevant.`,
		`The only difference between exceptional and amazing is you.`,
		`Shall I compare thee to a summer's day?  Thou art more lovely and more temperate.`,
		`there's bacon and then theres YOU`,
		`You make me think of beautiful things, like strawberries.`,
		`I would share my fruit Gushers with you.`,
		`You're more aesthetically pleasant to look at than that one green color on this website.`,
		`You're more fun than bubble wrap.`,
		`You make babies smile.`,
		`You make the gloomy days a little less gloomy.`,
		`You are warmer than a Snuggie.`,
		`You make me feel like I am on top of the world.`,
		`You remind me of my woobie`,
		`Let's never stop hanging out.`,
		`You're more cuddly than the Downy Bear.`,
		`You're so great I'd do your taxes`,
		`You are a foamy bucket of awesome.`,
		`If you really wanted to, you could probably get a bird to land on your shoulder and hang out with you.`,
		`My mom always asks me why I can't be more like you.`,
		`You know all the coolest music.`,
		`Chuck Norris told me he finds you midly intimidating.`,
		`Your body fat percentage is optimal with respect to your height.`,
		`I am having trouble coming up with a compliment worthy of you and I'm a COMPUTER.`,
		`If we were playing kickball, I'd pick you first.`,
		`You're cooler than ice on the rocks.`,
		`You're the bee's knees.`,
		`I wish I could choose your handwriting as a font.`,
		`You definitely know the difference between your and you're.`,
		`You have good taste.`,
		`You are grammatically superior.`,
		`I named all my appliances after you.`,
		`Don't worry about procrastinating on your studies, I know you'll do great!`,
		`I think about your for entire seconds (which is like 2e^6543212345 in computer years)`,
		`If you were in a chemistry class with me, it would be 10x less boring.`,
		`If you broke your arm, I would carry your books for you.`,
		`I love the way your eyes crinkle at the corners when you smile.`,
		`You make me want to be the person I am capable of being.`,
		`You're a skilled driver.`,
		`You are the rare catalyst to my volatile compound.`,
		`Looking at you makes my foot cramps go away instantaneously.`,
		`Cats like you.`,
		`You're so cool, that on a scale of from 1-10, you're elevendyseven.`,
		`You have the best laugh ever.`,
		`Your name is fun to say.`,
		`My camera isn't worthy to take your picture.`,
		`You are the sugar on my rice krispies.`,
		`You're real happening in a far out way, can you dig it?`,
		`Our awkward silences aren't even awkward.`,
		`I enjoy you more than a good sneeze. I mean like.. a GOOD one.`,
		`You could invent words and people would just *use* them.`,
		`You have powerful sweaters.`,
		`You are better than unicorns and sparkles combined!`,
		`You are the watermelon in my fruit salad. `,
		`I would trust my children with you.`,
		`You make me forget what I was going to...`,
		`I'd wake up for an 8 a.m. class just so I could sit next to you.`,
		`You have the moves like Jagger.`,
		`You're so hot that you denature my proteins.`,
		`All I want for Christmas is you!`,
		`You are the world's greatest hugger.`,
		`If you were a red shirt, they wouldn't kill you off.`,
		`They should name an ice cream flavor after you.`,
		`Me without you is like a nerd without braces, a shoe with out laces, asentencewithoutspaces.`,
		`Just knowing someone as cool as you will read this makes me smile.`,
		`I would volunteer to take your place in the Hunger Games.`,
		`If I had a thousand bucks for every time you did something stupid, I'd be broke!`,
		`I'd let you steal the white part of my Oreo.`,
		`The Force is strong with you.`,
		`I like the way your nostrils are placed on your nose.`,
		`I would hold the elevator doors open for you if they were closing.`,
		`You make me want to frolic in a field.`,
	}
	return fmt.Sprintf("Hey %s, %s",user, love[rand.Intn(len(love))])
}
