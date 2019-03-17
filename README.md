# instagen
Automatically generate instagram posts description from exif data.

## Setup
```bash
go get github.com/tomotetra/instagen
cd $GOPATH/src/github.com/tomotetra/instagen
make setup
```

## Usage
### Basic Usage
```bash
instagen path/to/image.jpg
```
This will generate the following:
```
2019.03.14
---
Canon EOS 5D Mark IV
EF24-70mm f/4L IS USM
ISO100 70mm f/8 1/80s
.
#東京カメラ部 #写真好きな人と繋がりたい #写真撮ってる人と繋がりたい #ファインダー越しの私の世界 #EXPLOREJPN #IG_JAPAN #IGersJP #PHOS_JAPAN #Rox_Captures #UnknownJapan #bestjapanpics #daily_photo_jpn #icu_japan #ig_myshot #instagramjapan #japan_of_insta #japantravelphoto #photo_shorttrip #retrip_news #team_jp_ #tokyocameraclub #visitjapanjp #wu_japan
```

You can set custom default hashtags by editing `$HOME/.instagenTags`.

**Warning**: Instagram currently limits users to 30 hashtags per post. Any hashtags exceeding this limit will be trimmed from the end.

### Options
| flag          | description                            | example                |
| ------------- | -------------------------------------- | ---------------------- |
| `--title, -T` | work title.                            | "a picure of a cat"    |
| `--tags, -t`  | additional hashtags separated by space | "cat kitten cats4life" |

```bash
instagen -T hoge -t "instagen foo bar" path/to/image.jpg
```
This will generate the following:
```
hoge
2019.03.14
---
Canon EOS 5D Mark IV
EF24-70mm f/4L IS USM
ISO100 70mm f/8 1/80s
.
#instagen #foo #bar #東京カメラ部 #写真好きな人と繋がりたい #写真撮ってる人と繋がりたい #ファインダー越しの私の世界 #EXPLOREJPN #IG_JAPAN #IGersJP #PHOS_JAPAN #Rox_Captures #UnknownJapan #bestjapanpics #daily_photo_jpn #icu_japan #ig_myshot #instagramjapan #japan_of_insta #japantravelphoto #photo_shorttrip #retrip_news #team_jp_ #tokyocameraclub #visitjapanjp #wu_japan
```
