# More on Decentralized Identity

### Why Does Identity Matter?

In the words of [Vitalik](https://vitalik.ca/general/2019/04/03/collusion.html):

>Mechanisms that do not rely on identity cannot solve the problem of concentrated interests outcompeting dispersed communities; an identity-free mechanism that empowers distributed communities cannot avoid over-empowering centralized plutocrats pretending to be distributed communities.

In other words, without an identity mechanism, one can't ensure "one human, one address" or "one human, one vote". This means that however you structure the rules of the system, those with the most resources would be able to game it.

### How Is the Existing System Failing Us?

In recent times, identities have been verified by credentials such as a passport or a social network account issued by a central authority (usually a state or corporation).

However, as noted in [Verifying Identity as a Social Intersection](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3375436), such identity systems have several interrelated flaws:

1. They are **insecure**. Crucial data such as an ID number constantly has to be given out. This is enough to impersonate an individual. On top of this, since all data is stored in a single repository managed by the state or a corporation, it becomes particularly vulnerable to hacking or internal corruption.

2. They **narrow** you down to one thing (in the system or out, a criminal or not, a credit score, etc.). The central database has little use for more information than this. This limits the functionality of the system and results in great injustices (for example, convicted individuals find it hard to re-enter society as this is the only information about them that they can reliably convey).

3. They are **artificial** in the sense that the information stored about you usually bears little relation to what you or your friends think of you about your identity.

To quote directly from the paper:

>Recently, new identity paradigms have tried to get around some of these elements. One approach, adopted by "big data" platforms like Facebook and Google, is to overcome thinness [narrowness] by storing enormous amounts of detailed information about each individual. we might call this "panoptic identity". However, such solutions have greatly exacerbated the other two problems, as they require extremely artificial compromises to intimacy through the global sharing of data with platforms that would not otherwise store it, creating exceptional potential security risks.

### Why Do We Need This Vision Now?

Given the rising political polarization and the increasing amount of information collected, shared, and cross-correlated by governments and corporations, there's a real risk that our information will be used against us in ways we cannot imagine.

If history has taught us anything, it's that **power belongs to those who control the information**.

Right now, that power belongs to the gatekeepers of our identities: governments and corporations.

In an increasingly uncertain world, there's a real risk that general fear, discontent and polarization will result in that power being abused.

In such a world, a check on the government and corporate power that goes beyond formal legal protection is essential.

By putting the control of information back in our hands, decentralized identity systems provide a natural technological check on the ability of the governments and corporations to abuse their power.

### How Can the Developing World Benefit?

In the developing world, decentralized identity systems have the potential to help bring millions of people out of the clutches of poverty.

To quote the words of [Timothy Ruff](https://medium.com/evernym/7-myths-of-self-sovereign-identity-67aea7416b1):

>Most of us take for granted that we can prove things about ourselves, unaware that over a billion people cannot. Identity is a prerequisite to financial inclusion, and financial inclusion is a big part of solving poverty.

### What Are Some of the Use Cases?

#### Liquid Democracy

Imagine if you could vote every two weeks to express your political sentiments regarding interest rates.

Imagine if you could have a direct say in any decision rather than relying on elected politicians to represent you. Imagine if those in power were held accountable in real-time rather than once every few years. This is the promise of liquid democracy. Liquid democracy exists somewhere in the sweet spot between direct and representative democracy.

As with direct democracy, everyone has the opportunity to vote on every issue. However, unlike direct democracy, you also have the choice to delegate your vote to someone else. You can even choose to delegate your votes on different issues to different people.

For example, on environmental issues, you might choose to delegate your vote to your favourite environmentalist. Whereas on issues concerning government debt and taxation you might choose your father.

This ability to delegate is recursive. It means that if your father, in turn, chooses to delegate his vote on financial issues to his favourite economist, your vote will also be delegated to that economist.

If you're unhappy with your father's decision, you can take that power away from him/her, and either vote yourself or re-delegate to someone you deem more trustworthy.

Under such a system, those with the most delegations become our representatives. But unlike representative democracy, they are held accountable in real-time.

>A system like this addresses the uninformed voters' issue that a direct democracy creates by allowing these voters allot their votes to experts. It also addresses the corruption issues of representative democracy because citizens can rescind their vote from someone instantly, forcing delegates to vote in the best interest of their constituents. It is the best of both worlds that truly gives the power of influence to the voters. [Source](https://media.consensys.net/liquid-democracy-and-emerging-governance-models-df8f3ce712af)

This sounds almost too good to be true. A fair, transparent, and corruption-free government! Why haven't we implemented this before?

Since there is no central government under this form of democracy, we need to figure out how to allow citizens to vote in a secure, private, and decentralized way. It turns out this is a pretty hard problem to solve. It has actually been impossible to solve. Until now. This is the first time in our history that technology exists to turn this dream into reality. Of course, we're talking about public blockchains.

Right now, we're in the experimentation phase. There are still some hard challenges that need to be overcome.

The three main challenges revolve around **scalability**, **privacy**, and **Sybil attacks**.

Scalability is important because we need millions of people to be able to use these systems.

Privacy is important because it ensures voters can't be discriminated against for the decisions they make. It also makes it harder for them to be bribed/coerced into voting for something they don't believe in.

But perhaps the hardest challenge is to ensure that one person is not able to vote multiple times (what's known in the jargon as a Sybil attack).

The key to solving the last two challenges is a voting protocol that requires basic verification and reputation for each user whilst protecting their pseudonymous identity. In other words, a voting protocol with a built-in decentralized identity system.

Put another way, decentralized identity is the big unlock that's needed to turn liquid democracy into a reality.

**P.S.** It turns out that solving the privacy problem helps solve the scalability problem too, but we won't get into that here.