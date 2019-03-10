# insta-template
Automatically generate instagram posts description from exif data.

## Usage
### Basic Usage
```bash
insta-template path/to/image.jpg
```
This will generate the following:
```
2019.03.10
---
Canon EOS 5D Mark IV
EF85mm f/1.4L IS USM
ISO125 85mm f/1.4 1/100s
.
#東京カメラ部 #広がり同盟 #写真好きな人と繋がりたい #写真撮ってる人と繋がりたい #ファインダー越しの私の世界 #EXPLOREJPN #ICU_VSCO #IG_JAPAN #IG_PHOS #pIGersJP #IGers_jp #LOVES_UNITED_JAPAN #Lovers_Nippon #PHOS_JAPAN #Photo_jpn #Rox_Captures #UnknownJapan #art_of_japan_ #bd_pro #bestjapanpics #bestphoto_japan #cityspride #daily_photo_jpn #icu_japan #ig_japan #ig_myshot #igworld_global #igworldglobal #instagramjapan #jalan_travel #japan_of_insta #japantravelphoto #jp_gallery #kf_gallery #photo_shorttrip #ptk_japan #retrip_news #s_shot #special_group__ #special_spot_ #super_japan_channel #team_jp_ #tokyocameraclub #visitjapanjp #whim_life #wu_japan
```

You can set custom default hashtags by editing [`common_hashtags.txt`](https://github.com/tomotetra/insta-template/blob/master/common_tags.txt).

### Options
| flag          | description                            | example                |
| ------------- | -------------------------------------- | ---------------------- |
| `--title, -T` | work title.                            | "a picure of a cat"    |
| `--tags, -t`  | additional hashtags separated by space | "cat kitten cats4life" |

```bash
insta-template -T hoge -t "foo bar fuga" path/to/image.jpg
```
This will generate the following:
```
2019.03.10
---
hoge
Canon EOS 5D Mark IV
EF85mm f/1.4L IS USM
ISO125 85mm f/1.4 1/100s
.
#foo #bar #fuga #東京カメラ部 #広がり同盟 #写真好きな人と繋がりたい #写真撮ってる人と繋がりたい #ファインダー越しの私の世界 #EXPLOREJPN #ICU_VSCO #IG_JAPAN #IG_PHOS #pIGersJP #IGers_jp #LOVES_UNITED_JAPAN #Lovers_Nippon #PHOS_JAPAN #Photo_jpn #Rox_Captures #UnknownJapan #art_of_japan_ #bd_pro #bestjapanpics #bestphoto_japan #cityspride #daily_photo_jpn #icu_japan #ig_japan #ig_myshot #igworld_global #igworldglobal #instagramjapan #jalan_travel #japan_of_insta #japantravelphoto #jp_gallery #kf_gallery #photo_shorttrip #ptk_japan #retrip_news #s_shot #special_group__ #special_spot_ #super_japan_channel #team_jp_ #tokyocameraclub #visitjapanjp #whim_life #wu_japan
```
