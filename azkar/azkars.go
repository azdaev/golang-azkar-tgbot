package azkar

import (
	"fmt"
	"strings"

	"github.com/azdaev/azkar-tg-bot/repository/models"
)

type Zikr struct {
	Arabic        string
	Russian       string
	Transcription string
}

var MorningAzkar = []Zikr{
	{`اللَّهُمَّ أَنْتَ رَبِّي لَا إِلَهَ إِلَّا أَنْتَ خَلَقْتَنِي وَ أَنَا عَبْدُكَ وَ أَنَا عَلَى عَهْدِكَ وَ وَعْدِكَ مَا اسْتَطَعْتُ أَعُوذُ بِكَ مِنْ شَرِّ مَا صَنَعْتُّ أَبُوءُ لَكَ بِنِعْمَتِكَ عَلَيَّ وَ أَبُوءُ بِذَنْبِي فَاغْفِرْ لِي فَإِنَّهُ لَا يَغْفِرُ الذُّنُوبَ إِلَّا أَنْتَ`,
		`О Аллах, Ты - Господь мой, и нет достойного поклонения, кроме Тебя, Ты создал меня, а я -Твой раб, и я буду хранить верность Тебе, пока у меня хватит сил. Прибегаю к Тебе от зла того, что я сделал, признаю милость, оказанную Тобой мне, и признаю грех свой. Прости же меня, ибо, поистине, никто не прощает грехов, кроме Тебя! (Бухари)`,
		`Аллахумма, Анта Рабби, ля иляха илля Анта, халякта-ни ва ана 'абду-кя, ва ана "аля 'ахди-кя ва ва'ди-кя ма-стата'ту. А'узу би-кя мин шарри ма сана'ту, абу'у ля-кя би-ни'мати-кя 'аляййя, ва абу'у би-занби, фа-гфир ли, фа-инна-ху ля йагфи-ру-з-зунуба илля Анта!`,
	},
	{`لَا إِلَهَ إِلَّا اللَّهُ وَحْدَهُ لَا شَرِيكَ لَهُ ، لَهُ الْمُلْكُ وَ لَهُ الْحَمْدُ وَ هُوَ عَلَى كُلِّ شَيْءٍ قَدِيرٌ`,
		`Нет достойного поклонения, кроме одного лишь Аллаха, у которого нет сотоварища, Ему принадлежит владычество, Ему хвала, Он всё может" 10 раз. (Ахмад)`,
		`Ля иляха илля-Ллаху вахда-ху ля шарикя ля-ху, ля-ху-ль-мульку ва ля-ху-ль-хамду ва хуа 'аля кулли шайин кадир.`,
	},
	{`سُبْحَانَ اللَّهِ وَ بِحَمْدِهِ`,
		`Пречист Аллах и хвала Ему. 100 раз. (Муслим)`,
		`Субхана-Ллахи ва би-хамди-хи.`,
	},
	{`أَصْبَحْنَا وَ أَصْبَحَ الْمُلْكُ لِلَّهِ ، وَ الْحَمْدُ لِلَّهِ ، لَا إِلَهَ إِلَّا اللَّهُ وَحْدَهُ لَا شَرِيكَ لَهُ ، لَهُ الْمُلْكُ وَ لَهُ الْحَمْدُ وَ هُوَ عَلَى كُلِّ شَيْءٍ قَدِيرٌ ، رَبِّ أَسْأَلُكَ خَيْرَ مَا فِي هَذَا الْيَوْمِ وَ خَيْرَ مَا بَعْدَهُ وَ أَعُوذُ بِكَ مِنْ شَرِّ مَا فِي هَذَا الْيَوْمِ وَ شَرِّ مَا بَعْدَهُ رَبِّ أَعُوذُ بِكَ مِنَ الْكَسَلِ وَ سُوءِ الْكِبَرِ ، رَبِّ أَعُوذُ بِكَ مِنْ عَذَابٍ فِي النَّارِ وَ عَذَابٍ فِي الْقَبْرِ`,
		`Мы дожили до утра, и этим утром владычество принадлежит Аллаху и хвала Аллаху, нет достойного поклонения, кроме одного лишь Аллаха, у которого нет сотоварища. Ему принадлежит владычество, Ему хвала, Он всё может. Господь мой, прошу Тебя о благе того, что будет в этот день, и благе того, что за ним последует, и прибегаю к Тебе от зла того, что будет в этот день, и зла того, что за ним последует. Господь мой, прибегаю к Тебе от нерадения и старческой дряхлости, Господь мой, прибегаю к Тебе от мучений в огне и мучений в могиле! (Муслим)`,
		`Асбахна ва асбаха-ль-мульку ли-Лляхи ва-ль-хамду ли-Лляхи, ля иляха илля Ллаху вахда-ху ля шарикя ля-ху, ля-ху-ль-мульку ва ля-ху-ль-хамду ва хуа ааля кулли шайин кадирун. Рабби ас'алюкя хайра ма фи хаза-ль-йауми ва хайра ма ба'да-ху ва а'узу би-кя мин шарри ма фи хаза-ль-йауми ва шарри ма ба'да-ху! Рабби, а'узу би-кя мин аль-кясали ва суи-ль-кибари, Рабби, а'узу би-кя мин 'азабин фи-н-нари ва 'азабин фи-ль-кабри!`,
	},
	{`اللَّهُمَّ بِكَ أَصْبَحْنَا وَبِكَ أَمْسَيْنَا وَبِكَ نَحْيَا وَبِكَ نَمُوتُ وَإِلَيْكَ النُّشُورُ`,
		`О Аллах, благодаря Тебе мы дожили до утра и благодаря Тебе мы дожили до вечера, Ты даёшь нам жизнь, и Ты лишаешь нас ее и Ты воскресишь нас для отчета. (Ахмад, абу Дауд, ат Тирмизи)`,
		`Аллахумма, би-кя асбахна, ва би-кя амсайна, ва би-кя нахйа, ва би-кя наму-ту ва иляй-кя-н-нушуру`,
	},
	{`بِسْمِ اللَّهِ الَّذِي لَا يَضُرُّ مَعَ اسْمِهِ شَيْءٌ فِي الْأَرْضِ وَ لَا فِي السَّمَاءِ وَ هُوَ السَّمِيعُ الْعَلِيمَُِ`,
		`С именем Аллаха, с именем которого ничто не причинит вред ни на земле, ни на небе, ведь Он - Слышащий, Знающий!" три раза. (Ахмад, ат Тирмизи)`,
		`Би-сми-Лляхи аллязи ля йадурру ма'а исми-хи шайун фи-ль-арди ва ля фи-с-самаи ва хуа-с-Сами'у-ль-'Алиму`,
	},
	{`اللَّهُمَّ إِنِّي أَسْأَلُكَ الْعَافِيَةَ فِي الدُّنْيَا وَ الْآخِرَةِ ، اللَّهُمَّ إِنِّي أَسْأَلُكَ الْعَفْوَ وَ الْعَافِيَةَ فِي دِينِي وَ دُنْيَايَ وَ أَهْلِي وَ مَالِي ، اللَّهُمَّ اسْتُرْ عَوْرَاتِي، وَآمِنْ رَوْعَاتِي ، اللَّهُمَّ احْفَظْنِي مِنْ بَيْنِ يَدَيَّ وَ مِنْ خَلْفِي وَ عَنْ يَمِينِي وَ عَنْ شِمَالِي وَ مِنْ فَوْقِي وَ أَعُوذُ بِعَظَمَتِكَ أَنْ أُغْتَالَ مِنْ تَحْتِي`,
		`O Аллах, поистине, я прошу Тебя о благополучии в мире этом и в мире ином, о Аллах, поистине, я прошу Тебя о прощении и благополучии в моей религии, и моих мирских делах, в моей семье и в моём имуществе. О Аллах, прикрой мою наготу и огради меня от страха, о Аллах, защити меня спереди, и сзади, и справа, и слева, и сверху, и я прибегаю к величию Твоему от того, чтобы быть предательски убитым снизу. (Ахмад, абу Дауд)`,
		`Аллахумма, инни ас'алю-кя-ль-'афийата фи-д-дунья ва-ль-ахирати, Аллахумма, инни ас'алю-кя-ль-'афуа ва-ль-'афийата фи дини, ва ду-ньяйа, ва ахли, ва мали. Аллахумма-стур 'аурати ва-эмин рау'ати, Аллахумма-хфаз-ни мин байни йадаййа, ва мин хальфи, ва 'ан ямини, ва 'ан шимали ва мин фауки, ва а'узу би-'азамати-кя ан угталя мин тахти!`,
	},
	{`اللَّهُمَّ عَالِمَ الْغَيْبِ وَ الشَّهَادَةِ ، فَاطِرَ السَّمَاوَاتِ وَ الْأَرْضِ ، رَبَّ كُلِّ شَيْءٍ وَ مَلِيكَهُ ، أَشْهَدُ أَنْ لَا إِلَهَ إِلَّا أَنْتَ أَعُوذُ بِكَ مِنْ شَرِّ نَفْسِي وَ مِنْ شَرِّ الشَّيْطَانِ وَ شِرْكِهِ وَ أَنْ أَقْتَرِفَ عَلَى نَفْسِي سُوءًا أَوْ أَجُرَّهُ إِلَى مُسْلِمٍ`,
		`O Аллах, Знающий сокрытое и явное, Творец небес и земли, Господь и Владыка всего, свидетельствую, что нет достойного поклонения, кроме Тебя, прибегаю к Тебе от зла души своей, от зла и многобожия шайтана и от того, чтобы причинить зло самому себе или навлечь его на какого-нибудь мусульманина. (Ахмад, ат Тирмизи)`,
		`Аллахумма, 'Алима-ль-гайби ва-ш-шахадати, Фатира-с-самавати ва-ль-арди, Рабба кулли шайин ва Малика-ху, ашхаду алля иляха илля Анта, а'узу би-кя мин шарри нафси, ва мин шарри-ш-шайтани ва ширки-хи ва ан актарифа 'аля нафси су'ан ау аджурра-ху иля мусли-мин.`,
	},
	{`أَصْبَحْنَا عَلَى فِطْرَةِ الْإِسْلَامِ وَ عَلَى كَلِمَةِ الْإِخْلَاصِ ، وَ عَلَى دِينِ نَبِيِّنَا مُحَمَّدٍ صَلَّى اللَّهُ عَلَيْهِ وَ سَلَّمَ وَ عَلَى مِلَّةِ أَبِينَا إِبْرَاهِيمَ حَنِيفاً مُسْلِماً وَ مَا كَانَ مِنَ الْمُشْرِكِينَ`,
		`Мы дожили до утра в лоне ислама согласно слову искренности, исповедуя религию нашего пророка Мухаммада, салляллаху ‘алейхи уа саллям, и религию нашего отца Ибрахима, который был ханифом и мусульманином и не относился к многобожникам. (Ахмад, ад Дарими).`,
		`Асбахна 'аля фитрати-ль-ислами ва 'аля кялимати-ль-ихляси ва 'аля дини набийй-на Мухаммадин, салля Ллаху 'аляй-хи ва салляма, ва 'аля мил-ляти аби-на Ибрахима ханифан муслиман ва ма кяна мин аль-мушрикина.`,
	},
	{`رَضِيتُ بِاللَّهِ رَبًّا ، وَ بِالْإسْلَامِ دِيناً وَ بِمُحَمَّدٍ صَلَّى اللَّهُ عَلَيْهِ وَسَلَّمَ نَبِيًّاَِ`,
		`Доволен я Аллахом как Господом, исламом - как религией и Мухаммадом, - как пророком!  Три раза. (Ахмад, абу Дауд)`,
		`Радийту би-Лляхи Раббан, ва би-ль-ислами динан ва би-Мухаммадин, сал-ля-Ллаху 'аляй-хи ва салляма, набийан`,
	},
	{`يَا حَيُّ يَا قَيُّومُ بِرَحْمَتِكَ أَسْتَغِيثُ أَصْلِحْ لِي شَأْنِي كُلَّهُ وَ لَا تَكِلْنِي إِلَى نَفْسِي طَرْفَةَ عَيْنٍ`,
		`О Живой, о Вечносущий, обращаюсь за защитой к милосердию Твоему, приведи в порядок все мои дела и не доверяй меня душе моей ни на мгновение! (ан Насаи).`,
		`Йа Хаййу, йа Кайюму, би-рахмати-кя астагису, аслих ли ша'ни кулля-ху ва ля такиль-ни иля нафси тарфата 'айнин!`,
	},
	{`سُبْحَانَ اللَّهِ وَ بِحَمْدِهِ عَدَدَ خَلْقِهِ وَ رِضَا نَفْسِهِ وَ زِنَةَ عَرشِهِ وَ مِدَادَ كَلِمَاتِهَِِ`,
		`Пречист Аллах и хвала Ему столько раз, сколько существует Его творений, и столько раз, сколько будет Ему угодно, пусть вес этих славословий и похвал будет равен весу Его трона и пусть для записи их потребуется столько же чернил, сколько нужно их для записи слов Его! три раза. (Муслим)`,
		`Субхана-Ллахи ва би-хамди-хи 'адада хальки-хи, ва рида нафси-хи, ва зината 'арши-хи ва мидада кялимати-хи!`,
	},
	{`حَسْبِيَ اللَّهُ لَا إِلَهَ إِلَّا هُوَ عَلَيْهِ تَوَكَّلْتُ وَ هُوَ رَبُّ الْعَرْشِ الْعَظِيمِ`,
		`Достаточно мне Аллаха, нет достойного поклонения, кроме Него, на Него я уповаю и Он - Господь великого трона. семь раз, (абу Дауд)`,
		`Хасбия-Ллаху, ля иляха илля хуа, 'аляй-хи таваккяльту ва хуа Раббу-ль-'арши-ль-'азыми`,
	},
}

var EveningAzkar = []Zikr{
	{`أَعُوذُ بِكَلِمَاتِ اللَّهِ التَّامَّاتِ مِنْ شَرِّ مَا خَلَقَ`,
		`Прибегаю к совершенным словам Аллаха от зла того, что Он сотворил. (Муслим)`,
		`А'узу би-кялимати Лляхи-т-таммати мин шарри ма халяка.`,
	},
	{`اللَّهُمَّ أَنْتَ رَبِّي لَا إِلَهَ إِلَّا أَنْتَ خَلَقْتَنِي وَ أَنَا عَبْدُكَ وَ أَنَا عَلَى عَهْدِكَ وَ وَعْدِكَ مَا اسْتَطَعْتُ أَعُوذُ بِكَ مِنْ شَرِّ مَا صَنَعْتُّ أَبُوءُ لَكَ بِنِعْمَتِكَ عَلَيَّ وَ أَبُوءُ بِذَنْبِي فَاغْفِرْ لِي فَإِنَّهُ لَا يَغْفِرُ الذُّنُوبَ إِلَّا أَنْتَ`,
		`О Аллах, Ты - Господь мой, и нет достойного поклонения, кроме Тебя, Ты создал меня, а я -Твой раб, и я буду хранить верность Тебе, пока у меня хватит сил. Прибегаю к Тебе от зла того, что я сделал, признаю милость, оказанную Тобой мне, и признаю грех свой. Прости же меня, ибо, поистине, никто не прощает грехов, кроме Тебя! (Бухари)`,
		`Аллахумма, Анта Рабби, ля иляха илля Анта, халякта-ни ва ана 'абду-кя, ва ана "аля 'ахди-кя ва ва'ди-кя ма-стата'ту. А'узу би-кя мин шарри ма сана'ту, абу'у ля-кя би-ни'мати-кя 'аляййя, ва абу'у би-занби, фа-гфир ли, фа-инна-ху ля йагфи-ру-з-зунуба илля Анта!`,
	},
	{`لَا إِلَهَ إِلَّا اللَّهُ وَحْدَهُ لَا شَرِيكَ لَهُ ، لَهُ الْمُلْكُ وَ لَهُ الْحَمْدُ وَ هُوَ عَلَى كُلِّ شَيْءٍ قَدِيرٌ`,
		`Нет достойного поклонения, кроме одного лишь Аллаха, у которого нет сотоварища, Ему принадлежит владычество, Ему хвала, Он всё может" 10 раз. (Ахмад)`,
		`Ля иляха илля-Ллаху вахда-ху ля шарикя ля-ху, ля-ху-ль-мульку ва ля-ху-ль-хамду ва хуа 'аля кулли шайин кадир.`,
	},
	{`سُبْحَانَ اللَّهِ وَ بِحَمْدِهِ`,
		`Пречист Аллах и хвала Ему. 100 раз. (Муслим)`,
		`Субхана-Ллахи ва би-хамди-хи.`,
	},
	{`أَمْسَيْنَا وَ أَمْسَى الْمُلْكُ لِلَّهِ ، وَ الْحَمْدُ لِلَّهِ ، لَا إِلَهَ إِلَّا اللَّهُ وَحْدَهُ لَا شَرِيكَ لَهُ ، لَهُ الْمُلْكُ وَ لَهُ الْحَمْدُ وَ هُوَ عَلَى كُلِّ شَيْءٍ قَدِيرٌ، رَبِّ أَسْأَلُكَ خَيْرَ مَا فِي هَذِهِ اللَّيْلَةِ وَ خَيْرَ مَا بَعْدَهَا وَ أَعُوذُ بِكَ مِنْ شَرِّ مَا فِي هَذِهِ اللَّيْلَةِ وَ شَرِّ مَا بَعْدَهَا رَبِّ أَعُوذُ بِكَ مِنَ الْكَسَلِ وَ سُوءِ الْكِبَرِ ، رَبِّ أَعُوذُ بِكَ مِنْ عَذَابٍ فِي النَّارِ وَ عَذَابٍ فِي الْقَبْرِ`,
		`Мы дожили до вечера, и этим вечером владычество принадлежит Аллаху, и хвала Аллаху, нет достойного поклонения, кроме одного лишь Аллаха, у которого нет сотоварища. Ему принадлежит владычество, Ему хвала, Он всё может. Господь мой, прошу Тебя о благе того, что будет в эту ночь, и благе того, что за ней последует, и прибегаю к Тебе от зла того, что будет в эту ночь, и зла того, что за ней последует. Господь мой, прибегаю к Тебе от нерадения и старческой дряхлости, Господь мой, прибегаю к Тебе от мучений в огне и мучений в могиле! (Муслим)`,
		`Амсайна ва амса-ль-мульку ли-Лляхи ва-ль-хамду ли-Лляхи, ля иляха илля Ллаху вахда-ху ля шарикя ля-ху, ля-ху-ль-мульку ва ля-ху-ль-хамду ва хуа ааля кулли шайин кадирун. Рабби ас'алюкя хайра ма фи хазихи-ль-лейляти ва хайра ма ба'адаха ва а'узу бикя мин шарри ма фи хазихи-ль-лейляти ва шарри ма ба'адаха Рабби, а'узу би-кя мин аль-кясали ва суи-ль-кибари, Рабби, а'узу би-кя мин 'азабин фи-н-нари ва 'азабин фи-ль-кабри!`,
	},
	{`اللَّهُمَّ بِكَ أَمْسَيْنَا وَبِكَ نَحْيَا وَبِكَ نَمُوتُ وَ إِلَيْكَ الْمَصِيرُ`,
		`О Аллах, благодаря Тебе мы дожили до вечера. Ты даёшь нам жизнь, и Ты лишаешь нас её и Ты воскресишь нас для отчета.`,
		`Аллахума, би-ка амсайна, ва би-ка нахйа, ва би-ка намуту ва иляй-ка- ль- масыру.`,
	},
	{`اللَّهُمَّ إِنِّي أَسْأَلُكَ الْعَافِيَةَ فِي الدُّنْيَا وَ الْآخِرَةِ ، اللَّهُمَّ إِنِّي أَسْأَلُكَ الْعَفْوَ وَ الْعَافِيَةَ فِي دِينِي وَ دُنْيَايَ وَ أَهْلِي وَ مَالِي ، اللَّهُمَّ اسْتُرْ عَوْرَاتِي، وَآمِنْ رَوْعَاتِي ، اللَّهُمَّ احْفَظْنِي مِنْ بَيْنِ يَدَيَّ وَ مِنْ خَلْفِي وَ عَنْ يَمِينِي وَ عَنْ شِمَالِي وَ مِنْ فَوْقِي وَ أَعُوذُ بِعَظَمَتِكَ أَنْ أُغْتَالَ مِنْ تَحْتِي`,
		`O Аллах, поистине, я прошу Тебя о благополучии в мире этом и в мире ином, о Аллах, поистине, я прошу Тебя о прощении и благополучии в моей религии, и моих мирских делах, в моей семье и в моём имуществе. О Аллах, прикрой мою наготу и огради меня от страха, о Аллах, защити меня спереди, и сзади, и справа, и слева, и сверху, и я прибегаю к величию Твоему от того, чтобы быть предательски убитым снизу. (Ахмад, абу Дауд)`,
		`Аллахумма, инни ас'алю-кя-ль-'афийата фи-д-дунья ва-ль-ахирати, Аллахумма, инни ас'алю-кя-ль-'афуа ва-ль-'афийата фи дини, ва ду-ньяйа, ва ахли, ва мали. Аллахумма-стур 'аурати ва-эмин рау'ати, Аллахумма-хфаз-ни мин байни йадаййа, ва мин хальфи, ва 'ан ямини, ва 'ан шимали ва мин фауки, ва а'узу би-'азамати-кя ан угталя мин тахти!`,
	},
	{`اللَّهُمَّ عَالِمَ الْغَيْبِ وَ الشَّهَادَةِ ، فَاطِرَ السَّمَاوَاتِ وَ الْأَرْضِ ، رَبَّ كُلِّ شَيْءٍ وَ مَلِيكَهُ ، أَشْهَدُ أَنْ لَا إِلَهَ إِلَّا أَنْتَ أَعُوذُ بِكَ مِنْ شَرِّ نَفْسِي وَ مِنْ شَرِّ الشَّيْطَانِ وَ شِرْكِهِ وَ أَنْ أَقْتَرِفَ عَلَى نَفْسِي سُوءًا أَوْ أَجُرَّهُ إِلَى مُسْلِمٍ`,
		`O Аллах, Знающий сокрытое и явное, Творец небес и земли, Господь и Владыка всего, свидетельствую, что нет достойного поклонения, кроме Тебя, прибегаю к Тебе от зла души своей, от зла и многобожия шайтана и от того, чтобы причинить зло самому себе или навлечь его на какого-нибудь мусульманина. (Ахмад, ат Тирмизи)`,
		`Аллахумма, 'Алима-ль-гайби ва-ш-шахадати, Фатира-с-самавати ва-ль-арди, Рабба кулли шайин ва Малика-ху, ашхаду алля иляха илля Анта, а'узу би-кя мин шарри нафси, ва мин шарри-ш-шайтани ва ширки-хи ва ан актарифа 'аля нафси су'ан ау аджурра-ху иля мусли-мин.`,
	},
	{`رَضِيتُ بِاللَّهِ رَبًّا ، وَ بِالْإسْلَامِ دِيناً وَ بِمُحَمَّدٍ صَلَّى اللَّهُ عَلَيْهِ وَسَلَّمَ نَبِيًّا`,
		`Доволен я Аллахом как Господом, исламом - как религией и Мухаммадом, - как пророком! Три раза. (Ахмад, абу Дауд)`,
		`Радийту би-Лляхи Раббан, ва би-ль-ислами динан ва би-Мухаммадин, сал-ля-Ллаху 'аляй-хи ва салляма, набийан`,
	},
	{`يَا حَيُّ يَا قَيُّومُ بِرَحْمَتِكَ أَسْتَغِيثُ أَصْلِحْ لِي شَأْنِي كُلَّهُ وَ لَا تَكِلْنِي إِلَى نَفْسِي طَرْفَةَ عَيْنٍ`,
		`О Живой, о Вечносущий, обращаюсь за защитой к милосердию Твоему, приведи в порядок все мои дела и не доверяй меня душе моей ни на мгновение! (ан Насаи)`,
		`Йа Хаййу, йа Кайюму, би-рахмати-кя астагису, аслих ли ша'ни кулля-ху ва ля такиль-ни иля нафси тарфата 'айнин!`,
	},
	{`حَسْبِيَ اللَّهُ لَا إِلَهَ إِلَّا هُوَ عَلَيْهِ تَوَكَّلْتُ وَ هُوَ رَبُّ الْعَرْشِ الْعَظِيمِ`,
		`Достаточно мне Аллаха, нет достойного поклонения, кроме Него, на Него я уповаю и Он - Господь великого трона. Семь раз. (абу Дауд)`,
		`Хасбия-Ллаху, ля иляха илля хуа, 'аляй-хи таваккяльту ва хуа Раббу-ль-'арши-ль-'азыми.`,
	},
	{
		`بِسْمِ اللَّهِ الَّذِي لَا يَضُرُّ مَعَ اسْمِهِ شَيْءٌ فِي الْأَرْضِ وَ لَا فِي السَّمَاءِ وَ هُوَ السَّمِيعُ الْعَلِيمُ`,
		`С именем Аллаха, с именем которого ничто не причинит вред ни на земле, ни на небе, ведь Он - Слышащий, Знающий!" три раза. (Ахмад, ат Тирмизи)`,
		`Би-сми-Лляхи аллязи ля йадурру ма'а исми-хи шайун фи-ль-арди ва ля фи-с-самаи ва хуа-с-Сами'у-ль-'Алиму`,
	},
}

func CurrentAzkarSlice(isMorning bool) []Zikr {
	if isMorning {
		return MorningAzkar
	}
	return EveningAzkar
}

func Wrap(config *models.ConfigInclude, index int, isMorning bool) string {
	var sb strings.Builder
	var zikr Zikr

	if isMorning {
		zikr = MorningAzkar[index]
		sb.WriteString("Утренний зикр ")
	} else {
		zikr = EveningAzkar[index]
		sb.WriteString("Вечерний зикр ")
	}

	sb.WriteString(fmt.Sprintf("№%d\n\n", index+1))

	if config.Arabic {
		sb.WriteString(zikr.Arabic)
		sb.WriteString("\n\n")
	}

	if config.Russian {
		sb.WriteString(zikr.Russian)
		sb.WriteString("\n\n")
	}

	if config.Transcription {
		sb.WriteString(zikr.Transcription)
	}

	return sb.String()
}
