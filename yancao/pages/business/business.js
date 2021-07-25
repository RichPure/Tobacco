// pages/business/business.js
Page({ 
  data: {
    items: [
      {id:1, value:'食杂店'},
      {id:2, value:'便利店'},
      {id:3, value:'烟酒商店'},
      {id:4, value:'超市商场'},
      {id:5, value:'娱乐服务业'},
    ],
    business:0
  },
  typeChange: function(e) {
    this.setData({
      business:e.detail.value
    })
    console.log('business：', this.data.business)
  },

  handleSubmit(e){
    console.log('submit business：', this.data.business)
    if (this.data.business == 0)
    {
      wx.showToast({
        title: '请选择！',
        icon: 'none'
      })
    }
    else if(this.data.business == 5)
    {
      wx.navigateTo({
        url: '/pages/room_num/room_num',
      })
    }
    else
    {
      wx.setStorageSync('business', this.data.business)
      wx.navigateTo({
        url: '/pages/grids/grids',
      })
    }
    
  }
})