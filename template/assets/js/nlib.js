(function () {
  var $ = function (selector) {
    return new $.fn.init(selector);
  };

  let sel;
  var getElement = function (selector) {
    sel=selector;
    if (typeof selector === "string") {
      return document.querySelector(selector);
    } else if (typeof selector.event == "object") {
      return selector.event.target;
    } else {
      return selector;
    }
  };

  $.fn = $.prototype = {
    init: function (selector) {
      //console.log(typeof selector, selector);
      this.elm = getElement(selector);
      return this;
    },
    sel:function(){
      return sel;
    },
    element: function () {
      return this.elm;
    },
    datas: function(key){
      if (typeof key==='undefined') return this.elm.dataset;
      return this.elm.dataset[key];
    },
    html: function (val) {
      if(typeof val==='undefined') return this.elm.innerHTML;
      this.elm.innerHTML=val;
      return this;
    },
    text: function (txt) {
      if (typeof txt === "string") {
        this.elm.innerText = txt;
        return this;
      } else {
        return this.elm.innerText;
      }
    },
    textc: function (txt) {
      if (txt) {
        this.elm.textContent = txt;
        return this;
      }
      return this.elm.textContent;
    },
    on: function (evtName, cb, options) {
      //console.log('on:',this.elm instanceof NodeList);
      if(this.elm instanceof NodeList){
        //this.elm.forEach(elm => elm.addEventListener(evtName, cb, options));
        for (let i = 0; i < this.elm.length; i++) {
            if (evtName=='click') this.elm[i].onclick=cb;
            else if (evtName=='change') this.elm[i].onchange=cb;
            else this.elm[i].addEventListener(evtName, cb, options);
        }
      }else{
        this.elm.addEventListener(evtName, cb, options);
      }
      return this;
    },
    click: function(cb){
      this.elm.onclick=cb;
      return this;
    },
    closest: function (sel) {
      this.elm = this.elm.closest(sel);
      return this;
    },
    children: function (cb) {
      //console.log(this.elm instanceof NodeList, this.elm instanceof HTMLCollection);
      this.elm = Array.from(this.elm.children); //HTMLCollection
      //console.log(this.elm);
      if (cb) this.elm.forEach(cb);
      return this;
    },
    childrenLength:function(){
      return this.elm.children.length;
    },
    child: function(index){
      if (typeof(index)==='number') this.elm=this.elm.childNodes[index];
      else this.elm=this.elm.childNodes;
      return this;
    },
    forms:function(){
      return this.elm.elements;
    },
    files:function(){
      return this.elm.files;
    },
    childItem: function (index) {
      //console.log(this.elm.children);
      this.elm = Array.from(this.elm.children)[index]; //HTMLCollection
      return this;
    },
    prev: function () {
      this.elm = this.elm.previousElementSibling;
      return this;
    },
    next: function () {
      this.elm = this.elm.nextElementSibling;
      return this;
    },
    parent: function () {
      this.elm = this.elm.parentElement;
      return this;
    },
    hide: function () {
      this.elm.style.display = "none";
      return this;
    },
    show: function () {
      this.elm.style.display = "";
      return this;
    },
    attr: function (name, value) {
      if (value == null) {
        return this.elm.getAttribute(name);
      } else {
        this.elm.setAttribute(name, value);
        return this;
      }
    },
    tag: function(){
      return this.elm.tagName;
    },
    prop: function(name, value){
      if (value == null) {
        return this.elm[name];
      }else{
        this.elm[name]=value;
        return this;
      }
    },
    val: function (s) {
      if (typeof s === "string") {
        this.elm.value = s;
        return this;
      } else {
        return this.elm.value;
      }
    },
    sindex: function () {
      return this.elm.options.selectedIndex;
    },
    index: function(){
      if(this.elm instanceof HTMLTableRowElement){
        return this.elm.rowIndex;
      }
      return this.elm.rowIndex;
    },
    stext: function () {
      return this.elm.options[this.sindex()].text;
    },
    find: function (sel) {
      //this.elm = this.elm.querySelector(sel); //HTMLElement
      let gls=this.elm.querySelectorAll(sel);
      if(gls.length==1){
        this.elm =gls[0];
      }else{
        this.elm =gls;
      }
      //console.log('find:',this.elm instanceof NodeList, this.elm instanceof HTMLCollection, this.elm instanceof HTMLElement);
      return this;
    },
    findAll: function (sel) {
      this.elm = this.elm.querySelectorAll(sel); //NodeList
      //console.log('findAll:',this.elm instanceof NodeList, this.elm instanceof HTMLCollection);
      return this;
    },
    isExist: function(){
      if (this.elm==null) return false;
      return this.elm.isConnected;
    },
    focus: function () {
      this.elm.focus();
      return this;
    },
    select: function () {
      this.elm.select();
      return this;
    },
    each: function (cb) {
      //console.log('each:',this.elm instanceof NodeList)
      //if (this.elm instanceof NodeList) 
      this.elm.forEach(cb); //NodeList
      return this;
    },
    isHidden:function(){
      return this.elm.style.display == "none" || this.elm.style.display == "";
    },
    nitem: function(index){
      if (this.elm instanceof NodeList) return this.elm[index];
    },
    matches: function (sel) {
      return this.elm.matches(sel);
    },
    is: function (sel) {
      return this.elm.matches(sel);
    },
    toggle: function (token) {
      this.elm.classList.toggle(token);
      return this;
    },
    contains: function (cls) {
      return this.elm.classList.contains(cls);
    },
    classList: function () {
      return this.elm.classList;
    },
    addClass: function (token) {
      if (Array.isArray(token)) {
        this.elm.classList.add(...token);
      } else {
        this.elm.classList.add(token);
      }
      return this;
    },
    append: function(){
      if (typeof arguments[0]==='string')
        this.elm.insertAdjacentHTML("beforeend", arguments[0]);
      else
      this.elm.append(...arguments);
      return this;
    },
    remove: function(){
      this.elm.remove();
      return this;
    },
    empty: function () {
      //this.elm.innerHTML = "";
      while (this.elm.hasChildNodes()) {
        this.elm.removeChild(this.elm.lastChild);
      }
      return this;
    },
    trigger: function (evtName) {
      let cvt = new Event(evtName);
      this.elm.dispatchEvent(cvt);
      return this;
    },
    removeClass: function (token) {
      if (Array.isArray(token)) {
        this.elm.classList.remove(...token);
      } else {
        this.elm.classList.remove(token);
      }
      return this;
    },
    submit: function (cb) {
      if (typeof cb==='undefined' && this.elm instanceof HTMLFormElement) this.elm.submit();
      //else this.elm.addEventListener("submit", cb);
      else this.elm.onsubmit=cb;
      return this;
    },
    serialize: function () {
      let arr = [];
      for (let i = 0; i < this.elm.elements.length; i++) {
        var element = this.elm.elements[i];
        let type = element.type;
        let kval = "";
        let key = encodeURIComponent(element.name);
        let val = encodeURIComponent(element.value);
        kval = `${key}=${val}`;
          if(type=="radio" || type=="checkbox"){
            if (element.checked){
              arr.push(kval);
            }
          }else if (type=="file"){
            //console.log(`${element.nodeName}-${element.nodeType}-${element.name}-${element.tagName}-${element.type}-${element.value}`);
          }else{
            //console.log(`${element.nodeName}-${element.nodeType}-${element.name}-${element.tagName}-${element.type}-${element.value}`);
            if (kval.length>2) arr.push(kval);
          }
      }
      let urlstr = arr.join("&");
      return urlstr;
    },
    formData: function () {
      var fdata = new FormData();
      for (let i = 0; i < this.elm.elements.length; i++) {
        var element = this.elm.elements[i];
        let etype = element.type;
        if (etype == "file") {
          let length=element.files.length;
          if (length==1){ 
            fdata.append(element.name, element.files[0], element.files[0].name); 
          }else{
            for (let i=0;i<length; i++){
              fdata.append(`${element.name}-${i}`, element.files[i], element.files[i].name);
            }
          }
        } else if (etype == "text") {
          fdata.append(element.name, element.value);
        } else if (etype == "radio") {
          if (element.checked) {
            fdata.append(element.name, element.value);
          }
        } else if (etype == "checkbox") {
          if (element.checked) {
            fdata.append(element.name, element.value);
          }
        } else if (etype == "textarea") {
          fdata.append(element.name, element.value);
        } else if (etype == "select") {
          fdata.append(element.name, element.value);
        }else{
          fdata.append(element.name, element.value);
        }
      }
      return fdata;
    },
    formSubmit: async function (url, fdata) {
      try {
        const res = await fetch(url, {
          method: "POST",
          mode: "cors",
          body: fdata,
        });
        const data = await res.json();
        if (!res.ok) {
          console.log(data.description);
          return;
        }
        return data;
      } catch (error) {
        console.log(error);
      }
    },
    postSubmit: async function (obj) {
      try {
        const rsp = await fetch(obj.url, {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
          },
          body: obj.data,
        });
        if (obj.json) {
          return await rsp.json();
        } else {
          return await rsp.text();
        }
      } catch (error) {
        console.log("ERR_" + error);
      }
    },
    fetch: function (obj) {
      if (typeof obj.data === "string") {
      }
      return fetch(obj.url, {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
        },
        body: obj.data,
      });
    },
    css: function (object) {
      for (let key in object) {
        this.elm.style[key] = object[key];
      }
      return this; //for chaining
    },
    removeAttr: function (attrName) {
      this.elm.removeAttribute(attrName);
      return this;
    },
    track: function (config, callback) {
      const observer = new MutationObserver(callback);
      observer.observe(this.elm, config);
      this.observer = observer;
      return this;
    },
    rect: function () {
      return this.elm.getBoundingClientRect();
    },
    top: function () {
      let rect = this.rect();
      return rect.top;
    },
    bottom: function () {
      let rect = this.rect();
      return window.innerHeight - rect.bottom;
    },
    insertElement: function (position, elment) {
      //??
      this.elm.insertAdjacentElement(position, elment);
      return this;
    },
    insertHTML: function (position, text) {
      //??
      this.elm.insertAdjacentHTML(position, text);
      return this;
    },
    insertAfter: function (element) {
      if (typeof element === "object") {
        this.insertElement("afterend", element);
      } else if (typeof element === "string") {
        this.insertHTML("afterend", element);
      }
      return this;
    },
    after: function(){
      this.elm.after(...arguments);
      return this;
    },
    crateAndInsert: function (strElm) {
      let element = this.createClassElement(strElm);
      this.elm.insertAdjacentElement("afterend", element);
      return this;
    },
    createClassElement: function (elementString) {
      let slc = elementString.split(".");
      const clsElm = document.createElement(slc[0]);
      for (let i = 1; i < slc.length; i++) {
        clsElm.classList.add(slc[i]);
      }
      return clsElm;
    },
    addChild: function (element) {
      this.elm.appendChild(element);
      return this;
    },
  };

  $.fn.init.prototype = $.fn;
  window.$ = $;
})();


$.ajax = function(obj){
  return new Promise( (resolve, reject) => {
      let body;
      const xhttp= new XMLHttpRequest();
      xhttp.open(obj.type, obj.url, obj.async);	
      if (obj.dataType){
         xhttp.responseType=obj.dataType;
      }
      if (obj.contentType=='application/json'){
          body=JSON.stringify(obj.data);
          xhttp.setRequestHeader("Content-type", "application/json");
      }else if (obj.data instanceof FormData) {
          body=obj.data;
      }else if (typeof obj.data=="object"){
          xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
          body=new URLSearchParams(obj.data).toString();
      }else{
          body=obj.data;
          xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      }
      xhttp.onreadystatechange = function(e) {
          if (xhttp.readyState == 4 && xhttp.status == 200) {
              if (obj.success) obj.success(xhttp.response);
          }else if(xhttp.readyState == 4 && xhttp.status != 200){
              if (obj.error) obj.error(xhttp.status, xhttp.response);
          }
      }
     xhttp.onload= ()=> {
          if(xhttp.status >= 200 && xhttp.status < 300){
              resolve(xhttp.response);
          }else{
              reject(`${xhttp.status}-${xhttp.response}`);
          }
     }
     xhttp.onerror= ()=> {
          reject(xhttp.response);
     }
     xhttp.send(body);
  });
}