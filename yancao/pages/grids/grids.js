// pages/grids/grids.js
Page({
  data: {
    items: [{
        id: 1,
        value: '繁华区'
      },
      {
        id: 2,
        value: '次繁华区'
      },
      {
        id: 3,
        value: '普通区'
      }
    ],
    grid: 0,
    grade: 0,
    latitude: 0,
    longitude: 0
  },

  bindGridInput(e) {
    this.setData({
      grid: e.detail.value
    })
  },

  handleChange(e) {
    this.setData({
      grade: e.detail.value
    })
    console.log("grade", this.data.grade)
  },

  handleGetLocation(e) {
    let that = this
    wx.getLocation({
      type: 'wgs84',
      //isHighAccuracy:true,
      success(res) {
        that.setData({
          latitude: res.latitude,
          longitude: res.longitude
        })
      }
    })
  },

  bindLatitudeInput(e) {
    this.setData({
      latitude: e.detail.value
    })
  },

  bindLongitudeInput(e) {
    this.setData({
      longitude: e.detail.value
    })
  },


  handleSubmit(e) {
    if (this.data.grid == 0 || this.data.latitude == 0 || this.data.longitude == 0) {
      wx.showToast({
        title: '请选择！',
        icon: 'none'
      })
    } else {
      wx.setStorageSync('grid', this.data.grid)
      wx.setStorageSync('grade', this.data.grade)
      wx.setStorageSync('latitude', this.data.latitude)
      wx.setStorageSync('longitude', this.data.longitude)
      let req_data = {
        'grid': this.data.grid,
        'grade': this.data.grade,
        'latitude': this.data.latitude,
        'longitude': this.data.longitude
      }
      console.log("req_data=", req_data)

      // resp_data = await request({
      //   url: "/inspection/get_grid ",
      //   data: req_data,
      //   method: "post"
      // });
      // console.log(resp_data);
      wx.setStorageSync('capacity', 10) //data.capacity)
      wx.setStorageSync('applied', 2)//data.applied)
      wx.setStorageSync('population', 3333)//data.population)
      wx.setStorageSync('distance', 88)//data.distance)      

      wx.navigateTo({
        url: '/pages/result/result',
      })
    }

  }
})