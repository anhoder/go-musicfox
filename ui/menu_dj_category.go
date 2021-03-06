package ui

import (
    "github.com/anhoder/netease-music/service"
    "go-musicfox/ds"
    "go-musicfox/utils"
)

type DjCategoryMenu struct {
    menus      []MenuItem
    categories []ds.DjCategory
}

func NewDjCategoryMenu() *DjCategoryMenu {
    return &DjCategoryMenu{}
}

func (m *DjCategoryMenu) MenuData() interface{} {
    return nil
}

func (m *DjCategoryMenu) BeforeBackMenuHook() Hook {
    return nil
}

func (m *DjCategoryMenu) IsPlayable() bool {
    return false
}

func (m *DjCategoryMenu) ResetPlaylistWhenPlay() bool {
    return false
}

func (m *DjCategoryMenu) GetMenuKey() string {
    return "dj_category"
}

func (m *DjCategoryMenu) MenuViews() []MenuItem {
    return m.menus
}

func (m *DjCategoryMenu) SubMenu(_ *NeteaseModel, index int) IMenu {
    if index >= len(m.categories) {
        return nil
    }

    return NewDjCategoryDetailMenu(m.categories[index].Id)
}

func (m *DjCategoryMenu) BeforePrePageHook() Hook {
    // Nothing to do
    return nil
}

func (m *DjCategoryMenu) BeforeNextPageHook() Hook {
    // Nothing to do
    return nil
}

func (m *DjCategoryMenu) BeforeEnterMenuHook() Hook {
    return func(model *NeteaseModel) bool {

        // 不重复请求
        if len(m.menus) > 0 && len(m.categories) > 0 {
            return true
        }

        djCateService := service.DjCatelistService{}
        code, response := djCateService.DjCatelist()
        codeType := utils.CheckCode(code)
        if codeType != utils.Success {
            return false
        }

        m.categories = utils.GetDjCategory(response)
        m.menus = GetViewFromDjCate(m.categories)

        return true
    }
}

func (m *DjCategoryMenu) BottomOutHook() Hook {
    return nil
}

func (m *DjCategoryMenu) TopOutHook() Hook {
    // Nothing to do
    return nil
}
