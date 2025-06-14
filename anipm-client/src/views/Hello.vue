<template>
  <el-card>
    <div style="margin-bottom: 14px; display: flex; align-items: center">
      <el-tag>帧: {{ currentFrame }}/{{ totalFrames }}</el-tag>
      <el-tag
        >时间: {{ (currentFrame * playInterval).toFixed(3) }}/{{
          (totalFrames * playInterval).toFixed(3)
        }}</el-tag
      >
      <el-tag type="info">帧率: {{ frameRate }}</el-tag>
      <el-tag type="info" style="margin-left: 12px"
        >每帧宽度: {{ frameBoxWidth }} px</el-tag
      >
      <el-button
        size="small"
        type="primary"
        style="margin-left: 16px"
        @click="togglePlay"
      >
        <span v-if="!isPlaying">▶️ 播放</span>
        <span v-else>⏸ 暂停</span>
      </el-button>
    </div>
    <el-scrollbar height="90">
      <div
        ref="timelineRef"
        class="frame-timeline"
        @wheel="onWheel"
        :style="{ width: containersWidth + 'px' }"
      >
        <!-- 帧格子 -->
        <div
          v-for="frame in frames"
          :key="frame"
          class="frame-box"
          :class="{ active: frame === currentFrame }"
          :style="{
            width: frameBoxWidth + 'px',
            height: frameBoxHeight + 'px',
          }"
          @mousedown="onFrameMouseDown(frame, $event)"
        >
          <span
            v-if="frame % frameNumberLabelStep === 1 && frameBoxWidth >= 32"
            class="frame-label"
            >{{ frame }}</span
          >
        </div>
        <!-- 指针 -->
        <div class="frame-pointer" :style="pointerStyle" />
      </div>
    </el-scrollbar>
  </el-card>
</template>

<script setup lang="ts">
import { ref, computed, nextTick, watch } from "vue";

const frameRate = 24;
const totalFrames = 1000;

const frameBoxMinWidth = 8;
const frameBoxMaxWidth = 64;
const frameBoxHeight = 50;
const frameNumberLabelStep = 10;

const frameBoxWidth = ref(16);
const frames = Array.from({ length: totalFrames }, (_, i) => i + 1);
const currentFrame = ref(1);
const timelineRef = ref<HTMLDivElement | null>(null);

const containersWidth = computed(() => totalFrames * frameBoxWidth.value + 2);

const pointerStyle = computed(() => ({
  transform: `translateX(${(currentFrame.value - 1) * frameBoxWidth.value}px)`,
  height: `${frameBoxHeight}px`,
  left: "0",
  width: "2px",
}));

let isDragging = false;
const onFrameMouseDown = (frame: number, _: MouseEvent) => {
  currentFrame.value = frame;
  isDragging = true;
  document.addEventListener("mousemove", onDrag);
  document.addEventListener("mouseup", offDrag);
  stopPlay();
};
const onDrag = (e: MouseEvent) => {
  if (!isDragging || !timelineRef.value) return;
  const timelineLeft = timelineRef.value.getBoundingClientRect().left;
  const x = Math.max(0, e.clientX - timelineLeft);
  let idx = Math.floor(x / frameBoxWidth.value) + 1;
  if (idx < 1) idx = 1;
  if (idx > totalFrames) idx = totalFrames;
  currentFrame.value = idx;
};
const offDrag = () => {
  isDragging = false;
  document.removeEventListener("mousemove", onDrag);
  document.removeEventListener("mouseup", offDrag);
};

function scrollIntoView(behavior: ScrollBehavior = "auto") {
  const frameDom = timelineRef.value?.children[
    currentFrame.value - 1
  ] as HTMLElement;
  frameDom?.scrollIntoView({
    behavior: behavior,
    block: "nearest",
    inline: "center",
  });
}

// 支持平滑缩放
function onWheel(e: WheelEvent) {
  if (e.shiftKey) return;
  e.preventDefault();
  let delta = e.deltaY > 0 ? 1 : -1;
  let step = Math.max(1, Math.floor(frameBoxWidth.value / 8));
  if (delta < 0) {
    frameBoxWidth.value = Math.min(
      frameBoxWidth.value + step,
      frameBoxMaxWidth
    );
  } else {
    frameBoxWidth.value = Math.max(
      frameBoxMinWidth,
      frameBoxWidth.value - step
    );
  }
  nextTick(() => scrollIntoView("instant"));
}

// 指针自动滚到可见区域
watch(currentFrame, () => {
  nextTick(() => scrollIntoView(isPlaying.value ? "instant" : "smooth"));
});

const isPlaying = ref(false);
let timer: number | null = null;
const playInterval = 1 / frameRate;

function togglePlay() {
  if (isPlaying.value) {
    stopPlay();
  } else {
    startPlay();
  }
}

function startPlay() {
  if (isPlaying.value) return;
  if (currentFrame.value >= totalFrames) {
    currentFrame.value = 1;
  }
  isPlaying.value = true;
  timer = window.setInterval(() => {
    if (currentFrame.value < totalFrames) {
      currentFrame.value++;
    } else {
      stopPlay();
    }
  }, playInterval);
}

function stopPlay() {
  isPlaying.value = false;
  if (timer) {
    clearInterval(timer);
    timer = null;
  }
}

// 离开页面、卸载组件前记得清理 timer
import { onUnmounted } from "vue";
onUnmounted(() => stopPlay());
</script>

<style scoped>
.frame-timeline {
  position: relative;
  min-width: 600px;
  display: flex;
  height: 50px;
  user-select: none;
  cursor: pointer;
  background: #fff;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
}
.frame-box {
  position: relative;
  height: 50px;
  box-sizing: border-box;
  border-right: 1px solid #dcdfe6;
  background: #f6f8fa;
  display: flex;
  align-items: flex-start;
  justify-content: center;
}
.frame-box.active {
  background: #cce2ff;
}
.frame-label {
  position: absolute;
  top: 28px;
  left: 2px;
  font-size: 11px;
  color: #aaa;
  user-select: none;
}
.frame-pointer {
  position: absolute;
  top: 0;
  z-index: 10;
  background: #409eff;
  width: 2px;
  border-radius: 1px;
  pointer-events: none;
  box-shadow: 0 0 4px #409eff55;
}
</style>
