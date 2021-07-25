import {
  request
} from "../../request/index.js";

Page({
  data: {
    capacity: 0,
    applied: 0,
    population: 0,
    distance: 0
  },

  onShow: function () {
    this.setData({
      capacity:wx.getStorageSync('capacity'),
      applied:wx.getStorageSync('applied'),
      population:wx.getStorageSync('population'),
      distance:wx.getStorageSync('distance')
    })
  },


  handleSubmit(e) {
    let req_data = {
      'business': wx.getStorageSync('business'),
      'room_num': wx.getStorageSync('room_num'),
      'grid': wx.getStorageSync('grid'),
      'grade': wx.getStorageSync('grade'),
      'latitude': wx.getStorageSync('latitude'),
      'longitude': wx.getStorageSync('longitude')
    }

    console.log("req_data=", req_data)
    // resp_data = await request({
    //   url: "/inspection/apply ",
    //   data: this.data.req_data,
    //   method: "post"
    // });
    wx.showModal({
      title: '实地勘验通过 ！',
      showCancel: false
    })
  }
})