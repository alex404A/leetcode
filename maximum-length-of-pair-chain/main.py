class Solution(object):
    def findLongestChain(self, pairs):
        """
        :type pairs: List[List[int]]
        :rtype: int
        """
        def sortMethod(p1, p2):
            if p1[0] > p2[0]:
                return 1
            elif p1[0] < p2[0]:
                return -1 
            if p1[0] == p2[0]:
                return p1[1] - p2[1]

        if len(pairs) == 0:
            return 0
        pairs = sorted(pairs, cmp=sortMethod)
        nonRepeatedPairs = [pairs[0]]
        pre = pairs[0]
        for i in range(1, len(pairs)):
            pair = pairs[i]
            if pair[0] != pre[0]:
                pre = pair
                for j in range(len(nonRepeatedPairs) - 1, -1, -1):
                    if nonRepeatedPairs[j][1] >= pair[1]:
                        nonRepeatedPairs.pop()
                    else:
                        break;
                nonRepeatedPairs.append(pair)
        print(nonRepeatedPairs)
        dp = [1] * len(nonRepeatedPairs)
        for i in range(len(nonRepeatedPairs)):
            for j in range(i):
                if nonRepeatedPairs[i][0] > nonRepeatedPairs[j][1]:
                    dp[i] = max(dp[i], dp[j] + 1)
        print(dp)
        return dp[len(nonRepeatedPairs) - 1]

if __name__ == '__main__':
    solution = Solution()
    pairs = [[-771,-143],[-293,578],[-404,-107],[-73,813],[427,793],[-141,67],[-537,43],[-744,246],[-371,238],[-31,19],[-906,-173],[273,817],[103,686],[553,971],[742,965],[-357,839],[662,762],[-726,876],[46,601],[-522,133],[21,449],[630,781],[-414,-302],[364,854],[871,965],[634,950],[-137,629],[21,864],[-133,929],[-876,-708],[313,518],[361,600],[-905,-11],[68,719],[492,822],[-67,68],[885,944],[920,954],[-988,-61],[698,967],[531,575],[961,966],[-554,56],[176,721],[316,647],[39,72],[707,801],[589,936],[-641,-473],[52,883],[12,345],[-226,163],[397,680],[666,727],[343,729],[442,449],[-363,-303],[524,840],[-742,692],[-180,147],[1,709],[304,912],[133,909],[-449,-282],[769,895],[-843,-327],[-742,60],[942,955],[555,715],[-838,-128],[-165,6],[-437,714],[-247,919],[-810,-478],[-36,72],[509,753],[338,561],[775,795],[236,898],[-636,205],[-114,343],[81,378],[748,749],[616,916],[240,884],[-671,-426],[-805,354],[-349,-315],[151,525],[448,603],[-193,151],[-370,-133],[904,913],[982,990],[938,956],[-586,-293],[988,996],[316,858],[902,971],[937,945],[356,552],[-689,842],[224,791],[196,943],[221,995],[-428,300],[776,933],[875,939],[-639,883],[-186,466],[-159,-105],[-653,156],[-872,-582],[506,992],[773,947],[754,930],[-496,-200],[699,709],[-495,420],[-12,762],[947,960],[-510,-302],[-244,100],[334,913],[505,947],[-549,-373],[-191,400],[32,209],[-481,921],[536,820],[-30,388],[131,251],[-248,920],[398,815],[846,881],[-204,79],[-678,-396],[116,980],[-566,245],[593,910],[-172,114],[-739,461],[882,904],[-229,-30],[862,899],[317,789],[226,695],[-708,-361],[-475,595],[261,324],[-399,-100],[-708,-354],[-409,795],[333,933],[-919,-158],[132,386],[-582,332],[33,949],[-177,265],[255,393],[76,658],[20,958],[728,815],[708,868],[54,349],[-713,386],[-535,295],[-177,286],[-341,556],[147,703],[419,564],[198,509],[-927,382],[548,875],[35,340],[-988,-230],[619,955],[-21,18],[608,830],[510,992],[-24,573],[-72,872],[-902,-871],[792,793],[-876,-624],[-832,-278],[162,806],[-600,-155],[-704,-411],[-90,764],[-565,740],[-451,202],[-223,272],[532,656],[958,998],[560,992],[-367,362],[-112,231],[987,996],[-348,311],[608,641],[211,514],[941,955],[-404,57],[6,180],[781,984],[-572,736],[341,710],[147,148],[-908,716],[-494,-342],[-209,870],[-900,-311],[53,935],[453,969],[-441,661],[-502,284],[-590,-420],[474,619],[-993,268],[-585,299],[859,893],[-667,-361],[571,840],[265,980],[33,490],[-830,-692],[-225,-126],[350,641],[346,708],[-412,835],[509,818],[-94,694],[71,657],[75,912],[-559,612],[804,916],[-806,390],[926,990],[438,588],[62,656],[164,576],[133,493],[-720,-424],[-13,103],[-624,436],[273,825],[-616,-13],[31,217],[240,339],[-775,826],[622,840],[-977,-535],[-761,-524],[-462,869],[608,707],[733,834],[373,958],[265,585],[325,807],[135,555],[-570,730],[-414,591],[-875,7],[486,736],[-922,-522],[-13,38],[-39,934],[963,983],[-13,968],[23,676],[913,926],[405,582],[-917,558],[-495,122],[908,968],[-831,-607],[977,980],[691,791],[-650,48],[211,393],[-489,987],[658,673],[-335,84],[536,843],[876,897],[975,988],[-195,471],[-713,-349],[-869,-724],[-162,270],[-77,-76],[314,952],[-620,-83],[157,361],[-544,770],[535,823],[-512,136],[897,916],[-830,775],[-13,218],[-195,-65],[-235,137],[701,913],[427,788],[939,989],[-484,-310],[-920,-134],[-366,125],[585,761],[-713,825],[-690,-365],[990,1000],[-57,135],[574,817],[226,848],[23,651],[-456,260],[-973,135],[-131,464],[-755,871],[-851,-668],[-469,-182],[236,546],[-99,879],[-180,695],[56,773],[-565,-401],[161,792],[-816,-782],[993,998],[34,852],[-255,819],[-733,949],[828,938],[-612,-177],[325,750],[333,951],[417,546],[-633,267],[995,997],[249,641],[351,803],[471,951],[-346,-259],[987,997],[834,872],[413,911],[-241,-54],[-276,869],[891,912],[349,836],[-684,-626],[753,829],[-970,-457],[-215,190],[-149,223],[-338,316],[-130,700],[391,552],[-495,922],[953,970],[577,645],[660,858],[148,512],[571,593],[-32,546],[222,311],[-50,800],[-911,330],[883,917],[612,754],[769,971],[703,791],[67,202],[-244,137],[-669,106],[60,804],[625,836],[-83,311],[-54,512],[-31,457],[-774,-407],[241,539],[-246,-11],[338,988],[875,970],[-398,-229],[-254,267],[-374,677],[-876,-863],[22,574],[-87,353],[857,899],[-398,910],[307,440],[-637,-375],[-940,-624],[343,620],[427,875],[-999,-232],[759,787],[496,832],[499,655],[-785,8],[370,837],[990,996],[-984,89],[-84,400],[285,292],[194,516],[-242,686],[915,943],[-261,625],[315,770],[-945,868],[947,990],[-983,646],[-806,76],[-567,-493],[998,1000],[446,847],[613,698],[199,672],[-661,-484],[296,898],[168,253],[-473,-207],[785,796],[533,609],[615,685],[623,723],[-463,838],[-298,896],[-67,6],[-533,959],[345,848],[983,996],[784,902],[-193,323],[912,961],[160,703],[664,790],[135,502],[-887,-645],[-490,917],[694,859],[-917,520],[201,254],[527,964],[-307,780],[-639,874],[863,879],[-729,-320],[-60,560],[-878,-280],[406,577],[543,910],[230,967],[-413,991],[-993,322],[856,978],[842,904],[120,998],[321,680],[500,583],[195,842],[416,467],[443,557],[977,984],[-929,435],[-382,605],[-543,866],[-507,399],[383,695],[868,954],[-284,246],[940,966],[-983,-50],[-457,3],[-357,74],[162,301],[289,496],[486,990],[-84,345],[-28,962],[-183,755],[-359,889],[-751,338],[-301,-249],[332,389],[498,970],[488,835],[305,416],[967,995],[932,998],[142,785],[417,844],[503,877],[407,620],[-611,726],[503,938],[119,388],[-355,190],[-88,490],[-148,33],[-47,153],[87,367],[379,449],[-895,79],[452,714],[-796,-483],[844,899],[-600,-227],[-555,-445],[-912,549],[129,943],[912,972],[425,508],[326,851],[403,960],[165,940],[-287,904],[112,700],[-813,-807],[285,619],[-131,787],[718,719],[603,609],[816,890],[308,564],[-423,155],[-408,32],[773,912],[752,933],[957,973],[-420,995],[-383,546],[503,881],[345,818],[-726,-613],[14,582],[-796,341],[-194,-99],[628,902],[-708,-229],[-721,-155],[138,596],[-890,-773],[782,893],[806,975],[-467,954],[691,884],[-363,361],[289,731],[580,835],[923,997],[83,90],[-275,808],[95,171],[203,658],[-89,615],[297,756],[129,598],[104,222],[410,975],[276,698],[-377,100],[-69,462],[145,207],[159,919],[-667,69],[-773,-567],[-588,-113],[599,717],[609,926],[624,684],[-713,998],[708,853],[411,549],[35,777],[-803,-287],[-601,-34],[-902,-554],[-553,456],[3,755],[649,974],[-439,290],[-393,515],[610,947],[76,543],[-50,289],[-348,122],[755,936],[-491,-66],[-360,112],[786,953],[-130,911],[-313,-34],[142,695],[-893,-780],[-433,244],[418,654],[707,927],[360,941],[-280,290],[-543,-60],[-982,136],[-95,154],[974,977],[573,810],[-850,-11],[-956,165],[761,805],[-514,-54],[825,976],[-703,753],[82,268],[717,919],[878,887],[-873,128],[267,740],[-169,896],[212,437],[-728,-345],[-697,460],[-128,558],[-257,826],[-191,139],[-858,432],[203,569],[195,755],[-128,462],[393,517],[-853,788],[752,775],[624,701],[661,754],[-16,475],[-336,-226],[-997,-142],[659,751],[-58,47],[756,867],[-513,828],[-786,989],[97,871],[952,954],[-852,-152],[-734,-715],[802,872],[-830,-218],[281,622],[-608,603],[768,870],[-763,681],[-844,953],[359,685],[181,547],[-785,437],[-857,488],[893,927],[-326,209],[-757,-148],[-78,364],[149,682],[-732,-408],[242,397],[-117,941],[-49,715],[142,337],[305,367],[-613,883],[638,760],[520,933],[-219,359],[275,697],[466,899],[-260,85],[-765,907],[644,981],[-785,-289],[-704,125],[-667,188],[91,733],[-548,138],[956,981],[411,660],[159,978],[-514,233],[-487,-270],[-727,159],[261,710],[949,968],[306,631],[-776,83],[382,773],[-867,-239],[249,418],[489,746],[-834,-546],[-938,573],[831,846],[880,963],[548,557],[-867,-476],[-535,746],[489,491],[771,802],[620,768],[377,498],[437,646],[-760,-612],[713,961],[622,678],[503,994],[-433,-226],[445,505],[743,939],[-401,749],[819,934],[12,837],[-965,-863],[-357,361],[412,1000],[-971,-116],[-588,230],[239,923],[-646,712],[58,593],[930,963],[331,730],[919,940],[139,792],[19,506],[660,923],[-258,738],[-117,852],[-528,380],[957,983],[132,321],[675,870],[-118,200],[-494,827],[437,729],[996,999],[169,847],[-846,188],[82,469],[-199,521],[245,866],[249,975],[-514,193],[-572,112],[362,739],[-492,557],[-681,-680],[-903,21],[111,500],[880,918],[-757,355],[-548,115],[-912,74],[686,763],[521,980],[-864,-18],[178,219],[-217,846],[887,983],[-603,754],[229,534],[-666,-589],[207,751],[-254,-110],[501,908],[-393,519],[883,973],[-138,55],[-957,-370],[-850,242],[-527,953],[-276,177],[985,988],[733,963],[-3,168],[992,998],[-932,-867],[9,508],[-233,-137],[999,1000],[-832,907],[786,914],[-266,144],[-570,-543],[-505,-343],[-397,45],[391,809],[321,923],[157,625],[-422,752],[101,894],[-290,-193],[-199,48],[869,988],[-975,800],[770,830],[-476,349],[-396,37],[-592,90],[-327,-117],[156,852],[464,919],[425,474],[547,853],[-685,-135],[-883,-743],[-912,-655],[879,957],[-859,-476],[289,924],[549,583],[539,729],[-891,8],[-337,700],[709,750],[839,896],[-147,527],[2,9],[894,914],[957,1000],[-899,895],[-966,341],[-163,-74],[-142,817],[917,988],[-719,-498],[-464,795],[354,964],[582,777],[823,863],[448,520],[-210,536],[-513,-79],[-262,639],[135,410],[-46,792],[-3,325],[318,339],[557,996],[872,919],[157,345],[502,813],[122,930],[-284,-116],[-29,911],[457,921],[420,541],[625,967],[538,630],[-153,667],[-814,-595],[-666,-117],[66,657],[62,483],[803,999],[667,966],[-421,-206],[687,891],[-451,-280],[-545,324],[950,993],[-719,-409],[-367,-289],[717,829],[-771,-280],[12,512],[-424,343],[-842,481],[-338,379],[-275,-91],[-616,947],[362,378],[-376,402],[811,830],[928,950],[-762,974],[131,339],[-347,888],[100,562],[485,829],[882,957],[830,947],[-882,394]]
    solution.findLongestChain(pairs)
