// pages/room_num/room_num.js
Page({
  data: {
    items: [{
        id: 5,
        value: '快餐饭店',
        input: '包间数量： '
      },
      {
        id: 6,
        value: '宾馆、KTV',
        input: '房间数量： '
      },
      {
        id: 7,
        value: '网吧',
        input: '散座数量： '
      },
    ],
    business: 0,
    room_num: 0,
  },

  typeChange: function (e) {
    this.setData({
      business: e.detail.value
    })
    console.log('business：', this.data.business)
  },

  bindKeyInput: function (e) {
    this.setData({
      room_num: e.detail.value
    })
    //console.log('get input：', this.data.room_num)
  },

  handleSubmit(e) {
    console.log('submit business：', this.data.business)
    if (this.data.business == 0 || this.data.room_num <= 0) {
      wx.showToast({
        title: '请选择并输入！',
        icon: 'none'
      })
    } else if ((this.data.business == 5 && this.data.room_num <= 10) ||
      (this.data.business == 6 && this.data.room_num <= 20) ||
      (this.data.business == 7 && this.data.room_num <= 80)) {
        wx.showToast({
          title: '未达到规模！',
          icon: 'none'
        }) 
    } else {
      wx.setStorageSync('business', this.data.business)
      wx.setStorageSync('room_num', this.data.room_num)
      wx.navigateTo({
        url: '/pages/grids/grids',
      })
    }

  }
})