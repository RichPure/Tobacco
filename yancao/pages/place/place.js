// pages/place/place.js
Page({
  data: {
    items: [
      {id:1, value:'幼儿园、中小学、未成年人教育机构周围'},
      {id:2, value:'安全隐患'},
      {id:3, value:'固定场所'},
      {id:4, value:'是否与住所独立'}
    ],
    ret_1:0,
    ret_2:0,
    ret_3:0,
    ret_4:0
  },

  handleChange(e) {

    let place_id = e.currentTarget.dataset.place_id;

    if(place_id == 1)
    {
      this.setData({
        ret_1:e.detail.value == "yes" ? 1 : 2
      })
    }
    else if(place_id == 2)
    {
      this.setData({
        ret_2:e.detail.value == "yes" ? 1 : 2
      })
    }
    else if(place_id == 3)
    {
      this.setData({
        ret_3:e.detail.value == "no" ? 1 : 2
      })
    }
    else
    {
      this.setData({
        ret_4:e.detail.value == "no" ? 1 : 2
      })
    }
    // console.log(place_id);
    // console.log(e.detail.value);
    // console.log("ret_1=" + this.data.ret_1+" ret_2=" + this.data.ret_2 + " ret_3=" + this.data.ret_3 + " ret_4=" + this.data.ret_4);
  },

  handleSubmit(e){
    // console.log("submit ret_1=" + this.data.ret_1+" ret_2=" + this.data.ret_2 + " ret_3=" + this.data.ret_3 + " ret_4=" + this.data.ret_4);
    if(this.data.ret_1 === 0 || this.data.ret_2 === 0 ||
      this.data.ret_3 === 0 || this.data.ret_4 === 0)
    {
      wx.showToast({
        title: '请全部勾选！',
        icon: 'none'
      })
    }
    else if(this.data.ret_1 === 2 && this.data.ret_2 === 2 &&
      this.data.ret_3 === 2 && this.data.ret_4 === 2)
    {
      wx.navigateTo({
        url: '/pages/ownership/ownership',
      })
    }
    else
    {
      wx.showToast({
        title: '不符合要求！',
        icon: 'none'
      })
    }
  }
})