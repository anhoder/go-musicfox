package ui

import (
    "fmt"
    "github.com/anhoder/netease-music/service"
    "go-musicfox/ds"
    "go-musicfox/utils"
    "strconv"
)

type DjRadioDetailMenu struct {
    menus     []MenuItem
    songs     []ds.Song
    djRadioId int64
    limit     int
    offset    int
}

func NewDjRadioDetailMenu(djRadioId int64) *DjRadioDetailMenu {
    return &DjRadioDetailMenu{
        djRadioId: djRadioId,
        limit:     50,
        offset:    0,
    }
}

func (m *DjRadioDetailMenu) MenuData() interface{} {
    return m.songs
}

func (m *DjRadioDetailMenu) BeforeBackMenuHook() Hook {
    return nil
}

func (m *DjRadioDetailMenu) IsPlayable() bool {
    return true
}

func (m *DjRadioDetailMenu) ResetPlaylistWhenPlay() bool {
    return false
}

func (m *DjRadioDetailMenu) GetMenuKey() string {
    return fmt.Sprintf("dj_radio_detail_%d", m.djRadioId)
}

func (m *DjRadioDetailMenu) MenuViews() []MenuItem {
    return m.menus
}

func (m *DjRadioDetailMenu) SubMenu(_ *NeteaseModel, _ int) IMenu {
    return nil
}

func (m *DjRadioDetailMenu) BeforePrePageHook() Hook {
    // Nothing to do
    return nil
}

func (m *DjRadioDetailMenu) BeforeNextPageHook() Hook {
    // Nothing to do
    return nil
}

func (m *DjRadioDetailMenu) BeforeEnterMenuHook() Hook {
    return func(model *NeteaseModel) bool {

        djProgramService := service.DjProgramService{
            RID:    strconv.FormatInt(m.djRadioId, 10),
            Limit:  strconv.Itoa(m.limit),
            Offset: strconv.Itoa(m.offset),
        }
        code, response := djProgramService.DjProgram()
        codeType := utils.CheckCode(code)
        if codeType != utils.Success {
            return false
        }
        m.songs = utils.GetSongsOfDjRadio(response)
        m.menus = GetViewFromSongs(m.songs)

        return true
    }
}

func (m *DjRadioDetailMenu) BottomOutHook() Hook {
    // Nothing to do
    return nil
}

func (m *DjRadioDetailMenu) TopOutHook() Hook {
    // Nothing to do
    return nil
}
