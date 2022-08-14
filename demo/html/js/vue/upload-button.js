Vue.component('z-upload-button', {
  template: '<div class="z-upload-button" style="position: relative">' +
      '<div class="zu-button"><a class="uploadImg">' +
        '<input type="file" accept="*/*" @change="upload" ref="iUpload"  />' +
        '<span>{{label}}</span></a></div>' +
      '<div class="zu-input" ><el-input v-model="value" ></el-input></div>' +
    '</div>',
  props :['label','value','base64','item'],
  methods : {
    upload(e){
      let _this = this;
      let fileList = e.target.files;
      for(let i=0; i< fileList.length; i++){
        let name = fileList[i]["name"];
        if (this.base64){
          getImageBase64(fileList[i], function(data){
            data = new File([data], name, {
              type: 'text/plain;charset=utf-8'
            });
            _this.$emit("change",{e:e,item:_this.item,name, data});
          });
        }else{
          _this.$emit("change",{e:e,item:_this.item,name, data: fileList[i]});
        }
        
        
      }
    }
  }
});