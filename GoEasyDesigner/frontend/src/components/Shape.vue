<template>
  <div class="Shape" ref="shapeRef">
    <slot></slot>
    <div
        v-if="index == nowIndex"
        v-for="(direction, index) in directions"
        :key="index"
        :class="['dot', direction]"
        :style="direction"
        @mousedown="mousedown(direction, $event)"
    ></div>
  </div>
</template>

<script>
export default {
  name: "DirectionalDots",
  props: {
    index: {
      type: Number,
      default: 0,
    },
    nowIndex: {
      type: Number,
      default: 0,
    },
  },
  created() {
    console.log("index", this.index);
    console.log("nowIndex", this.nowIndex);

  },
  emits: ['update-style'],
  data() {
    return {
      directions: ["top", "bottom", "left", "right", "top-left", "top-right", "bottom-left", "bottom-right"],
      draggingDot: "",
      clientX: 0,
      clientY: 0,
      startX: 0,
      startY: 0,
    };
  },
  methods: {
    mousedown: function (direction, event) {
      event.stopPropagation()
      event.preventDefault()
      this.draggingDot = direction;
      this.startX = event.clientX;
      this.startY = event.clientY;
      let isFirst = true;
      let dthis = this;

      let newLeft2 = dthis.$refs.shapeRef.offsetLeft;
      let newTop2 = dthis.$refs.shapeRef.offsetTop;
      let newWidth2 = dthis.$refs.shapeRef.offsetWidth;
      let newHeight2 = dthis.$refs.shapeRef.offsetHeight;

      let mousemove = function (event) {
        // console.log("移动", dthis.startX, dthis.startY, event.clientX, event.clientY);

        dthis.clientX = event.clientX;
        dthis.clientY = event.clientY;
        if (isFirst) {
          isFirst = false
          return
        }
        //计算移动距离
        let moveX = dthis.clientX - dthis.startX;
        let moveY = dthis.clientY - dthis.startY;

        // console.log("放开",  dthis.startX, dthis.startY,dthis.clientX, dthis.clientY);
        // console.log("移动距离",dthis.draggingDot, moveX, moveY);

        let newLeft = newLeft2;
        let newTop = newTop2;
        let newWidth = newWidth2;
        let newHeight = newHeight2;

        if (dthis.draggingDot === "top") {
          newTop += moveY;
          newHeight -= moveY;
        } else if (dthis.draggingDot === "bottom") {
          newHeight += moveY;
        } else if (dthis.draggingDot === "left") {
          newLeft += moveX;
          newWidth -= moveX;
        } else if (dthis.draggingDot === "right") {
          newWidth += moveX;
        } else if (dthis.draggingDot === "top-left") {
          newLeft += moveX;
          newWidth -= moveX;
          newTop += moveY;
          newHeight -= moveY;
        } else if (dthis.draggingDot === "top-right") {
          newWidth += moveX;
          newTop += moveY;
          newHeight -= moveY;
        } else if (dthis.draggingDot === "bottom-left") {
          newLeft += moveX;
          newWidth -= moveX;
          newHeight += moveY;
        } else if (dthis.draggingDot === "bottom-right") {
          newWidth += moveX;
          newHeight += moveY;
        }
        if (newWidth < 0 || newHeight < 0) {
          return
        }
        // console.log("移动后的位置",  dthis.index,newLeft, newTop, newWidth, newHeight);
        dthis.$emit('update-style', dthis.index,{
          left: newLeft,
          top: newTop,
          width: newWidth,
          height: newHeight
        })
      }
      let mouseup = function () {
        document.removeEventListener("mousemove", mousemove);
        document.removeEventListener("mouseup", mouseup);
      }
      document.addEventListener("mousemove", mousemove);
      document.addEventListener("mouseup", mouseup);

    },

  },
};
</script>

<style scoped>
.Shape {
  position: absolute;
}

.dot {
  width: 6px;
  height: 6px;
  background-color: #057aff;
  position: absolute;
  cursor: pointer;
}

.top {
  top: -6px;
  left: 50%;
  transform: translateX(-50%);
  cursor: n-resize;
}

.bottom {
  bottom: -8px;
  left: 50%;
  transform: translateX(-50%);
  cursor: s-resize;
}

.left {
  top: 50%;
  left: -6px;
  transform: translateY(-50%);
  cursor: w-resize;
}

.right {
  top: 50%;
  right: -8px;
  transform: translateY(-50%);
  cursor: e-resize;
}

.top-left {
  top: -6px;
  left: -6px;
  cursor: nw-resize;
}

.top-right {
  top: -6px;
  right: -8px;
  cursor: ne-resize;
}

.bottom-left {
  bottom: -6px;
  left: -6px;
  cursor: sw-resize;
}

.bottom-right {
  bottom: -8px;
  right: -8px;
  cursor: se-resize;
}
</style>