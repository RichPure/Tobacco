// pages/ownership/ownership.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    items: [{
        id: 1,
        value: '无偿使用'
      },
      {
        id: 2,
        value: '自有'
      },
      {
        id: 3,
        value: '租赁'
      }
    ],
    ownership: 0,
    rent: 0
  },

  handleRent(e) {
    this.setData({
      rent: e.detail.value == "yes" ? 1 : 2
    })
  },


  typeChange: function (e) {
    this.setData({
      ownership: e.detail.value
    })
    console.log('ownership', this.data.ownership)
  },

  handleSubmit(e) {
    console.log('submit ownership', this.data.ownership)
    if (this.data.ownership == 0) {
      wx.showToast({
        title: '请选择！',
        icon: 'none'
      })
      return
    }

    if (this.data.ownership == 3) {
      if (this.data.rent == 0) {
        wx.showToast({
          title: '请选择租期！',
          icon: 'none'
        })
        return
      } else if (this.data.rent == 2) {
        wx.showToast({
          title: '租期应大于1年！',
          icon: 'none'
        })
        return
      }
    }

    wx.setStorageSync('ownership', this.data.ownership)
    wx.navigateTo({
      url: '/pages/business/business',
    })
  }
})