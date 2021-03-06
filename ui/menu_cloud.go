package ui

import (
    "github.com/anhoder/netease-music/service"
    "github.com/buger/jsonparser"
    "go-musicfox/ds"
    "go-musicfox/utils"
    "strconv"
)

type CloudMenu struct {
    menus   []MenuItem
    songs   []ds.Song
    limit   int
    offset  int
    hasMore bool
}

func NewCloudMenu() *CloudMenu {
    return &CloudMenu{
        limit:  100,
        offset: 0,
    }
}

func (m *CloudMenu) MenuData() interface{} {
    return m.songs
}

func (m *CloudMenu) BeforeBackMenuHook() Hook {
    return nil
}

func (m *CloudMenu) IsPlayable() bool {
    return true
}

func (m *CloudMenu) ResetPlaylistWhenPlay() bool {
    return false
}

func (m *CloudMenu) GetMenuKey() string {
    return "could"
}

func (m *CloudMenu) MenuViews() []MenuItem {
    return m.menus
}

func (m *CloudMenu) SubMenu(*NeteaseModel, int) IMenu {
    return nil
}

func (m *CloudMenu) BeforePrePageHook() Hook {
    // Nothing to do
    return nil
}

func (m *CloudMenu) BeforeNextPageHook() Hook {
    // Nothing to do
    return nil
}

func (m *CloudMenu) BeforeEnterMenuHook() Hook {
    return func(model *NeteaseModel) bool {
        if utils.CheckUserInfo(model.user) == utils.NeedLogin {
            NeedLoginHandle(model, enterMenu)
            return false
        }

        // 不重复请求
        if len(m.menus) > 0 && len(m.songs) > 0 {
            return true
        }

        cloudService := service.UserCloudService{
            Offset: strconv.Itoa(m.offset),
            Limit:  strconv.Itoa(m.limit),
        }
        code, response := cloudService.UserCloud()
        codeType := utils.CheckCode(code)
        if codeType == utils.NeedLogin {
            NeedLoginHandle(model, enterMenu)
            return false
        } else if codeType != utils.Success {
            return false
        }

        if hasMore, err := jsonparser.GetBoolean(response, "hasMore"); err == nil {
            m.hasMore = hasMore
        }

        m.songs = utils.GetSongsOfCloud(response)
        m.menus = GetViewFromSongs(m.songs)

        return true
    }
}

func (m *CloudMenu) BottomOutHook() Hook {
    if !m.hasMore {
        return nil
    }
    return func(model *NeteaseModel) bool {
        m.offset += m.limit

        // 不重复请求
        if len(m.menus) > 0 && len(m.songs) > 0 {
            return true
        }

        cloudService := service.UserCloudService{
            Offset: strconv.Itoa(m.offset),
            Limit:  strconv.Itoa(m.limit),
        }
        code, response := cloudService.UserCloud()
        codeType := utils.CheckCode(code)
        if codeType == utils.NeedLogin {
            NeedLoginHandle(model, enterMenu)
            return false
        } else if codeType != utils.Success {
            return false
        }

        if hasMore, err := jsonparser.GetBoolean(response, "hasMore"); err == nil {
            m.hasMore = hasMore
        }

        songs := utils.GetSongsOfCloud(response)
        menus := GetViewFromSongs(songs)

        m.songs = append(m.songs, songs...)
        m.menus = append(m.menus, menus...)

        return true
    }
}

func (m *CloudMenu) TopOutHook() Hook {
    // Nothing to do
    return nil
}
