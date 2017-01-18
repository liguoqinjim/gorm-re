package main

type Test2 struct {
	Id  int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Tid int `gorm:"column:tid;not null"`
}

type CampAhSoldier struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	CasSoId      int    `gorm:"column:cas_so_id;not null"`
	CasNum       int    `gorm:"column:cas_num;not null"`
	CasPriceType int    `gorm:"column:cas_price_type;not null"`
	CasPrice     int    `gorm:"column:cas_price;not null"`
	CasPlayerId  int    `gorm:"column:cas_player_id;not null"`
	CasTotalTime int    `gorm:"column:cas_total_time;not null"`
	CasStartTime string `gorm:"column:cas_start_time;size:30;not null"`
	CasIsSell    int    `gorm:"column:cas_isSell;not null;default:0"`
}

type CityInfo struct {
	Id             int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId       int `gorm:"column:player_id;not null"`
	RegionId       int `gorm:"column:region_id;not null"`
	RegionSecondId int `gorm:"column:region_second_id;not null"`
	CityX          int `gorm:"column:city_x;not null"`
	CityY          int `gorm:"column:city_y;not null"`
	CityTime       int `gorm:"column:city_time;not null"`
}

type EmbassyHelpOther struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	HelpId   int `gorm:"column:help_id;not null"`
}

type EmbassySeekHelp struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	HelpType int `gorm:"column:help_type;not null"`
	ThingId  int `gorm:"column:thing_id;not null"`
	GuildId  int `gorm:"column:guild_id;not null"`
	MaxTimes int `gorm:"column:max_times;not null"`
	NowTimes int `gorm:"column:now_times;not null"`
	Help     int `gorm:"column:help;not null"`
}

type GuildApply struct {
	Id            int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildId       int    `gorm:"column:guild_id;not null"`
	PlayerId      int    `gorm:"column:player_id;not null"`
	PlayerName    string `gorm:"column:player_name;size:16;not null"`
	ApplyTime     int    `gorm:"column:apply_time;not null"`
	WhetherHandle int    `gorm:"column:whether_handle;not null;default:0"`
}

type GuildFlag struct {
	Id           int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildId      int `gorm:"column:guild_id;not null"`
	FlagId       int `gorm:"column:flag_id;not null"`
	PurchaserId  int `gorm:"column:purchaser_id;not null"`
	PurchaseTime int `gorm:"column:purchase_time;not null"`
}

type GuildGoods struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildId  int `gorm:"column:guild_id;not null"`
	GoodsId  int `gorm:"column:goods_id;not null"`
	GoodsCnt int `gorm:"column:goods_cnt;not null"`
}

type GuildGoodsMessage struct {
	Id          int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildId     int `gorm:"column:guild_id;not null"`
	MessageType int `gorm:"column:message_type;not null"`
	PlayerId    int `gorm:"column:player_id;not null"`
	PlayerAuth  int `gorm:"column:player_auth;not null"`
	GoodsId     int `gorm:"column:goods_id;not null"`
	GoodsCnt    int `gorm:"column:goods_cnt;not null"`
	MessageTime int `gorm:"column:message_time;not null"`
}

type GuildInfo struct {
	Id                int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildName         string `gorm:"column:guild_name;size:64;not null"`
	GuildAbbr         string `gorm:"column:guild_abbr;size:16;not null"`
	GuildCreatorId    int    `gorm:"column:guild_creator_id;not null"`
	GuildCreatorName  string `gorm:"column:guild_creator_name;size:16;not null"`
	GuildCreateTime   int    `gorm:"column:guild_create_time;not null"`
	GuildLevel        int    `gorm:"column:guild_level;not null"`
	GuildPower        int    `gorm:"column:guild_power;not null"`
	GuildRace         int    `gorm:"column:guild_race;not null"`
	GuildMemberCount  int    `gorm:"column:guild_member_count;not null"`
	GuildHonor        int    `gorm:"column:guild_honor;not null;default:0"`
	GuildAllowApply   int    `gorm:"column:guild_allow_apply;not null"`
	GuildApplyFactor1 int    `gorm:"column:guild_apply_factor1;not null"`
	GuildApplyFactor2 int    `gorm:"column:guild_apply_factor2;not null"`
	GuildDeclaration  string `gorm:"column:guild_declaration;size:128;not null"`
	GuildFlag         int    `gorm:"column:guild_flag;not null"`
	GuildFlagColor    string `gorm:"column:guild_flag_color;size:6;not null"`
	GuildEmblem       int    `gorm:"column:guild_emblem;not null"`
	GuildEmblemColor  string `gorm:"column:guild_emblem_color;size:6;not null"`
	GuildCountry      int    `gorm:"column:guild_country;not null"`
	GuildLanguage     int    `gorm:"column:guild_language;not null"`
	GuildClass        string `gorm:"column:guild_class;size:64;not null"`
	GuildOrg          string `gorm:"column:guild_org;size:64;not null"`
}

type GuildMember struct {
	Id                int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildId           int `gorm:"column:guild_id;not null"`
	PlayerId          int `gorm:"column:player_id;not null"`
	GuildAuth         int `gorm:"column:guild_auth;not null"`
	GuildContribution int `gorm:"column:guild_contribution;not null"`
}

type GuildMessage struct {
	Id          int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildId     int    `gorm:"column:guild_id;not null"`
	SenderId    int    `gorm:"column:sender_id;not null"`
	SenderAuth  int    `gorm:"column:sender_auth;not null"`
	MessageType int    `gorm:"column:message_type;not null"`
	Message     string `gorm:"column:message;size:64;not null"`
	SendTime    int    `gorm:"column:send_time;not null"`
}

type GuildTechnology struct {
	Id         int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildId    int `gorm:"column:guild_id;not null"`
	TechId     int `gorm:"column:tech_id;not null"`
	TechLv     int `gorm:"column:tech_lv;not null"`
	TechExp    int `gorm:"column:tech_exp;not null"`
	FinishTime int `gorm:"column:finish_time;not null"`
	DonateType int `gorm:"column:donate_type;not null"`
}

type GuildTechnologyDonate struct {
	Id         int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId   int `gorm:"column:player_id;not null"`
	GuildId    int `gorm:"column:guild_id;not null"`
	TechId     int `gorm:"column:tech_id;not null"`
	DonateType int `gorm:"column:donate_type;not null"`
}

type GuildWholeMessage struct {
	Id          int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	GuildId     int    `gorm:"column:guild_id;not null"`
	SenderId    int    `gorm:"column:sender_id;not null"`
	MessageType int    `gorm:"column:message_type;not null"`
	Message     string `gorm:"column:message;size:64;not null"`
	SendTime    int    `gorm:"column:send_time;not null"`
}

type Player struct {
	PlayerId       int    `gorm:"column:Player_id;primary_key;AUTO_INCREMENT;not null"`
	PlayerName     string `gorm:"column:Player_name;size:20;not null"`
	PlayerPassword string `gorm:"column:Player_password;size:100;not null"`
	PlayerRace     int    `gorm:"column:Player_race;default:0"`
	PlayerLv       int    `gorm:"column:Player_lv;default:1"`
}

type Player9Formation struct {
	Id        int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	FId       int    `gorm:"column:f_id;not null"`
	PlayerId  int    `gorm:"column:Player_id;not null"`
	Formation string `gorm:"column:formation;size:128"`
}

type Player9RacingTask struct {
	Id               int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	TId              int    `gorm:"column:t_id;not null"`
	TBid             int    `gorm:"column:t_bid;not null"`
	TStart           int    `gorm:"column:t_start;not null"`
	TFirstFinish     int    `gorm:"column:t_first_finish;not null"`
	TFirstFinishName string `gorm:"column:t_first_finish_name;size:32;not null"`
	TFinishNum       int    `gorm:"column:t_finish_num;not null"`
	TMapId           int    `gorm:"column:t_map_id;not null"`
}

type Player9RacingTaskProgress struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	Task1    int `gorm:"column:task_1;not null"`
	Task2    int `gorm:"column:task_2;not null"`
	Task3    int `gorm:"column:task_3;not null"`
	TStart   int `gorm:"column:t_start;not null"`
	MapId    int `gorm:"column:map_id;not null;default:0"`
}

type Player9Task struct {
	Id         int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	TId        int    `gorm:"column:t_id;not null"`
	PlayerId   int    `gorm:"column:player_id;not null"`
	TbId       int    `gorm:"column:tb_id;not null"`
	TStar      int    `gorm:"column:t_star;not null"`
	TStartTime int    `gorm:"column:t_start_time;not null"`
	TFinish    int    `gorm:"column:t_finish;not null"`
	TCond      string `gorm:"column:t_cond;size:64;not null"`
}

type Player9TaskProgress struct {
	Id         int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId   int `gorm:"column:player_id;not null"`
	Progress1  int `gorm:"column:progress_1;not null"`
	Progress2  int `gorm:"column:progress_2;not null"`
	Progress3  int `gorm:"column:progress_3;not null"`
	Progress4  int `gorm:"column:progress_4;not null"`
	Progress5  int `gorm:"column:progress_5;not null"`
	Progress6  int `gorm:"column:progress_6;not null"`
	Progress7  int `gorm:"column:progress_7;not null"`
	Progress8  int `gorm:"column:progress_8;not null"`
	Progress9  int `gorm:"column:progress_9;not null"`
	Progress10 int `gorm:"column:progress_10;not null"`
	Progress11 int `gorm:"column:progress_11;not null"`
	Progress12 int `gorm:"column:progress_12;not null"`
	Progress13 int `gorm:"column:progress_13;not null"`
	Progress14 int `gorm:"column:progress_14;not null"`
	Progress15 int `gorm:"column:progress_15;not null"`
	Progress16 int `gorm:"column:progress_16;not null"`
}

type PlayerArmy struct {
	AId        int    `gorm:"column:A_id;primary_key;AUTO_INCREMENT;not null"`
	APlayerId  int    `gorm:"column:A_player_id;not null"`
	AHeroId    int    `gorm:"column:A_hero_id;not null"`
	AMapId     int    `gorm:"column:A_map_id;not null"`
	ARegionId  int    `gorm:"column:A_region_id;not null"`
	AX         int    `gorm:"column:A_x"`
	AY         int    `gorm:"column:A_y"`
	ASoldier   string `gorm:"column:A_soldier;size:64"`
	AFormation string `gorm:"column:A_formation;size:128"`
	ADirection int    `gorm:"column:A_direction"`
	ACurrent   string `gorm:"column:A_current;size:100"`
	APaths     string `gorm:"column:A_paths;size:1000"`
}

type PlayerBond struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	BondId   int `gorm:"column:bond_id;not null"`
}

type PlayerCity struct {
	Id          int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId    int `gorm:"column:player_id;not null"`
	BuildingId  int `gorm:"column:building_id;not null"`
	BuildingLv  int `gorm:"column:building_lv;not null"`
	BuildingPos int `gorm:"column:building_pos;not null"`
	FinishTime  int `gorm:"column:finish_time;not null"`
	Time1       int `gorm:"column:time_1;not null"`
	Time2       int `gorm:"column:time_2;not null"`
	Factor1     int `gorm:"column:factor_1;not null"`
	Factor2     int `gorm:"column:factor_2;not null"`
}

type PlayerEquipment struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	EId      int `gorm:"column:e_id;not null"`
	HId      int `gorm:"column:h_id;not null"`
	EPos     int `gorm:"column:e_pos;not null"`
}

type PlayerEquipmentForge struct {
	Id          int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId    int    `gorm:"column:player_id;not null"`
	ELv         int    `gorm:"column:e_lv;not null"`
	Equipments  string `gorm:"column:equipments;size:128;not null"`
	EIndex      int    `gorm:"column:e_index;not null"`
	ForgeTimes  int    `gorm:"column:forge_times;not null"`
	ForgeTimes2 int    `gorm:"column:forge_times2;not null"`
}

type PlayerEvent struct {
	EventId      int `gorm:"column:event_id;primary_key;not null"`
	EventVersion int `gorm:"column:event_version;not null"`
	PlayerId     int `gorm:"column:player_id;primary_key;not null"`
}

type PlayerGoods struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	GoodsId  int `gorm:"column:goods_id;not null"`
}

type PlayerGov struct {
	Id                int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId          int    `gorm:"column:player_id;not null"`
	PlayerCgId        int    `gorm:"column:player_cg_id;not null"`
	PlayerDailyGet    int    `gorm:"column:player_daily_get;not null"`
	PlayerFlowerTimes int    `gorm:"column:player_flower_times;not null"`
	PlayerModelName   string `gorm:"column:player_model_name;size:20;not null"`
	PlayerIsGetSalary int    `gorm:"column:player_isGet_salary;not null"`
}

type PlayerGovFormation struct {
	Id               int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId         int    `gorm:"column:player_id;not null"`
	PlayerFormations string `gorm:"column:player_formations;size:200;not null"`
}

type PlayerGovGoldSoldier struct {
	Id         int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId   int `gorm:"column:player_id;not null"`
	SoldierId  int `gorm:"column:soldier_id;not null"`
	SoldierNum int `gorm:"column:soldier_num;not null"`
}

type PlayerGovMission struct {
	Id          int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId    int `gorm:"column:player_id;not null"`
	CgmId       int `gorm:"column:cgm_id;not null"`
	PcmPassTime int `gorm:"column:pcm_pass_time;not null"`
}

type PlayerGovNor struct {
	Id                   int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId             int    `gorm:"column:player_id;not null"`
	PlayerCgnId          int    `gorm:"column:player_cgn_id;not null"`
	PlayerModelName      string `gorm:"column:player_model_name;size:20;not null"`
	PlayerIsGetSalary    int    `gorm:"column:player_isGet_salary;not null"`
	PlayerCampRegisterId int    `gorm:"column:player_camp_register_id;not null"`
}

type PlayerGovRegister struct {
	CgId       int    `gorm:"column:cg_id;primary_key;not null"`
	CgRegister string `gorm:"column:cg_register;size:500;not null"`
}

type PlayerGovSoldier struct {
	Id         int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId   int `gorm:"column:player_id;not null"`
	SoldierId  int `gorm:"column:soldier_id;not null"`
	SoldierNum int `gorm:"column:soldier_num;not null"`
}

type PlayerHero struct {
	HId       int `gorm:"column:H_id;primary_key;AUTO_INCREMENT;not null"`
	HPlayerId int `gorm:"column:H_player_id;not null"`
	HBaId     int `gorm:"column:H_ba_id;not null"`
	HLv       int `gorm:"column:H_lv;not null"`
	HExp      int `gorm:"column:H_exp;not null"`
	HStar     int `gorm:"column:H_star;not null"`
	HSkillLv1 int `gorm:"column:H_skill_lv1;not null"`
	HSkillLv2 int `gorm:"column:H_skill_lv2;not null"`
	HSkillLv3 int `gorm:"column:H_skill_lv3;not null"`
	HSkillLv4 int `gorm:"column:H_skill_lv4;not null"`
}

type PlayerHeroOwned struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	HeroId   int `gorm:"column:hero_id;not null"`
	HeroCnt  int `gorm:"column:hero_cnt;not null"`
}

type PlayerItem struct {
	PIId       int `gorm:"column:PI_id;primary_key;AUTO_INCREMENT;not null"`
	PIPlayerId int `gorm:"column:PI_player_id"`
	PIMapId    int `gorm:"column:PI_map_id"`
	PIItmId    int `gorm:"column:PI_itm_id"`
	PINum      int `gorm:"column:PI_num"`
}

type PlayerRegimentSkill struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	SkillId  int `gorm:"column:skill_id;not null"`
	HeroId   int `gorm:"column:hero_id;not null"`
}

type PlayerResource struct {
	PlayerId     int `gorm:"column:player_id;primary_key;AUTO_INCREMENT;not null"`
	PCityX       int `gorm:"column:p_city_x;not null;default:0"`
	PCityY       int `gorm:"column:p_city_y;not null;default:0"`
	PGold        int `gorm:"column:p_gold;not null;default:0"`
	PDiamond     int `gorm:"column:p_diamond;not null;default:0"`
	PWood        int `gorm:"column:p_wood;not null;default:0"`
	PWood2       int `gorm:"column:p_wood2;not null;default:0"`
	PFood        int `gorm:"column:p_food;not null;default:0"`
	PFood2       int `gorm:"column:p_food2;not null;default:0"`
	PIron        int `gorm:"column:p_iron;not null;default:0"`
	PIron2       int `gorm:"column:p_iron2;not null;default:0"`
	PSteel       int `gorm:"column:p_steel;not null;default:0"`
	PSteel2      int `gorm:"column:p_steel2;not null;default:0"`
	PGuild       int `gorm:"column:p_guild;not null;default:0"`
	PGuildHonor  int `gorm:"column:p_guildHonor;not null;default:0"`
	PFlowerTimes int `gorm:"column:p_flower_times;not null;default:5"`
	PGovPoint    int `gorm:"column:p_gov_point;not null;default:0"`
}

type PlayerSkill struct {
	SId       int `gorm:"column:S_id;primary_key;not null;default:0"`
	SMapId    int `gorm:"column:S_map_id;primary_key;not null;default:0"`
	SPlayerId int `gorm:"column:S_player_id;primary_key;not null;default:0"`
	SNum      int `gorm:"column:S_num"`
}

type PlayerSoldier struct {
	Id         int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId   int `gorm:"column:Player_id;default:0"`
	SoldierId  int `gorm:"column:Soldier_id;default:0"`
	SoldierNum int `gorm:"column:Soldier_num;default:0"`
}

type PlayerSoldierUnlock struct {
	Id        int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId  int `gorm:"column:player_id;not null"`
	SoldierId int `gorm:"column:soldier_id;not null"`
}

type PlayerTechnology struct {
	Id           int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId     int `gorm:"column:player_id;not null"`
	TechId       int `gorm:"column:tech_id;not null"`
	TechTreeType int `gorm:"column:tech_tree_type;not null"`
	TechLv       int `gorm:"column:tech_lv;not null"`
	FinishTime   int `gorm:"column:finish_time;not null"`
}

type PlayerTopGoods struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	GoodsId  int `gorm:"column:goods_id;not null"`
	NowCnt   int `gorm:"column:now_cnt;not null"`
	Appeared int `gorm:"column:appeared;not null"`
}

type PlayerTreasure struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	TId      int `gorm:"column:t_id;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	SId1     int `gorm:"column:s_id1;not null"`
	SId2     int `gorm:"column:s_id2;not null"`
	SId3     int `gorm:"column:s_id3;not null"`
	SFactor1 int `gorm:"column:s_factor1;not null"`
	SFactor2 int `gorm:"column:s_factor2;not null"`
	SFactor3 int `gorm:"column:s_factor3;not null"`
	SGrade1  int `gorm:"column:s_grade1;not null"`
	SGrade2  int `gorm:"column:s_grade2;not null"`
	SGrade3  int `gorm:"column:s_grade3;not null"`
	TPos     int `gorm:"column:t_pos;not null"`
}

type PlayerTreasurePos struct {
	Id       int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId int `gorm:"column:player_id;not null"`
	Pos      int `gorm:"column:pos;not null"`
}

type PlayerTreasureRecast struct {
	Id       int    `gorm:"column:id;primary_key;not null"`
	PlayerId int    `gorm:"column:player_id;not null"`
	SId1     int    `gorm:"column:s_id1;not null"`
	SId2     int    `gorm:"column:s_id2;not null"`
	SId3     int    `gorm:"column:s_id3;not null"`
	SFactor1 int    `gorm:"column:s_factor1;not null"`
	SFactor2 int    `gorm:"column:s_factor2;not null"`
	SFactor3 int    `gorm:"column:s_factor3;not null"`
	SGrade1  int    `gorm:"column:s_grade1;not null"`
	SGrade2  int    `gorm:"column:s_grade2;not null"`
	SGrade3  int    `gorm:"column:s_grade3;not null"`
	SLocks   string `gorm:"column:s_locks;size:3;not null"`
}

type PlayerTreasureTask struct {
	Id             int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId       int `gorm:"column:player_id;not null"`
	TId            int `gorm:"column:t_id;not null"`
	TStartTime     int `gorm:"column:t_start_time;not null"`
	THid1          int `gorm:"column:t_hid1;not null"`
	THid2          int `gorm:"column:t_hid2;not null"`
	THid3          int `gorm:"column:t_hid3;not null"`
	THero1         int `gorm:"column:t_hero1;not null"`
	THero2         int `gorm:"column:t_hero2;not null"`
	THero3         int `gorm:"column:t_hero3;not null"`
	TSug           int `gorm:"column:t_sug;not null"`
	TNeed          int `gorm:"column:t_need;not null"`
	TPrize         int `gorm:"column:t_prize;not null"`
	TSuggestFactor int `gorm:"column:t_suggest_factor;not null"`
	TNeedFactor    int `gorm:"column:t_need_factor;not null"`
}

type PlayerTypeEvent struct {
	Id          int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	PlayerId    int `gorm:"column:player_id;not null"`
	TypeEventId int `gorm:"column:type_event_id;not null"`
	Cond1       int `gorm:"column:cond_1;not null"`
	Cond2       int `gorm:"column:cond_2;not null"`
	Value1      int `gorm:"column:value_1;not null"`
	Value2      int `gorm:"column:value_2;not null"`
}

type RegionInfo struct {
	Id             int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	RegionId       int `gorm:"column:region_id;not null"`
	RegionSecondId int `gorm:"column:region_second_id;not null"`
	CityCnt        int `gorm:"column:city_cnt;not null"`
}

type ScheduleEvent struct {
	EventId       int    `gorm:"column:event_id;primary_key;not null"`
	EventVersion  int    `gorm:"column:event_version;not null"`
	EventLastTime string `gorm:"column:event_last_time;size:40;not null"`
}

type Server struct {
	SId     int    `gorm:"column:S_id;primary_key;not null"`
	SName   string `gorm:"column:S_name;size:20;not null"`
	SIp     string `gorm:"column:S_ip;size:20;not null"`
	SPort   string `gorm:"column:S_port;size:20;not null"`
	SStatus int    `gorm:"column:S_status;not null"`
}
