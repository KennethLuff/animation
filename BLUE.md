>### DATABASE
>* user:
>	>id, nickname, head_icon, weikey, head, address, account, tel, password, 
user_city, user_city_code
>* user-video:
>	>id, video
>* video:
>	>id, title, ...model, music, type, status, rack
>* model:
>	>id, type, title, result
>* shared:
>	>id, video, uid, fee
>* music:
>	>id, uid, type
>* base-music:
>* ad:
>	>id, image, time
>* type:


>### URL
>* login:
>* ytpe:
>* user:
>	>user/info
>	>user/video
>* index:

>### admin function url
>* admin/addSlenes
>* admin/addTemp
>* admin/ad

>### MODEL
>* Loag:
>	> Loagin
>		> OtherLoagin
>	> LoagOut
>* Account:
>	> Register
>	> UserInfo
>	> UserInfoUpdate
>	> UserVideo
>	> ResatePassword
>* Shared:
>	> Shared
>* Ad:
>	> ControlAd
>* util:
>	 > ComposeVideo
>* pay:
>	> wx.pay	

>### Dir
>* config: 存放项目的配置文件
>* dal:	存放与数据库交互相关代码
>* logic:	存放业务逻辑相关代码
>* model:	存放项目的数据结构相关代码
>* static:	存放项目的静态资源
>* util:	存放项目的一些工具功能以及杂项
>* views:	存放项目的一些工具功能以及杂项
>* controller:	存放控制器