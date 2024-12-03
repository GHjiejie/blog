<template>
  <div class="comment-container">
    <div class="comment-header">
      <div class="comment-header-title">评论:0</div>
    </div>
    <div class="comment-content">
      <div class="comment-content-left">
        <div class="avatar">
          <img
            src="https://images.pexels.com/photos/957040/night-photograph-starry-sky-night-sky-star-957040.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2"
            alt=""
          />
        </div>
      </div>
      <div class="comment-content-right">
        <div class="comment-content-right-header">
          <div class="comment-content-right-header-left">
            <div class="comment-content-right-header-left-username">用户名</div>
            <div class="comment-content-right-header-left-time">2021-10-10</div>
          </div>
          <div class="comment-content-right-header-right">
            <div class="comment-content-right-header-right-like">
              <i class="icon"></i>
            </div>
          </div>
        </div>
        <div class="comment-content-right-center">
          <p>
            CTFShow通常是指网络安全领域中的“Capture The
            Flag”(夺旗赛)展示工具或平台。这是一种用于分享、学习和展示信息安全竞赛中获取的信息、漏洞利用技巧以及解题思路的在线社区或软件。参与者会在比赛中收集“flag”，通常是隐藏在网络环境中的数据或密码形式，然后通过分析、破解等手段找到并提交。CTFShow可以帮助人们了解最新的安全技术和挑战，同时也促进了安全知识和技术的交流。
          </p>
        </div>
      </div>
    </div>
    <div class="comment-footer">
      <div class="comment-footer-input">
        <input type="text" placeholder="请输入评论内容" />
      </div>
      <div class="comment-footer-button">
        <button>评论</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineExpose } from "vue";
</script>

<style scoped lang="scss">
.comment-container {
  width: 100%;
  margin: 0 auto;
  border: 1px solid #ccc;
  border-radius: 5px;
  overflow: hidden;
  position: relative;

  .comment-header {
    padding: 10px;
    background-color: #f0f0f0;
    border-bottom: 1px solid #ccc;

    .comment-header-title {
      font-size: 16px;
      font-weight: bold;
    }
  }

  .comment-content {
    display: flex;
    padding: 10px;

    .comment-content-left {
      width: 50px;
      margin-right: 10px;

      .avatar {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        overflow: hidden;

        img {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
      }
    }

    .comment-content-right {
      flex: 1;

      .comment-content-right-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 5px;

        .comment-content-right-header-left {
          display: flex;
          align-items: center;

          .comment-content-right-header-left-username {
            font-size: 14px;
            font-weight: bold;
            margin-right: 10px;
          }

          .comment-content-right-header-left-time {
            font-size: 12px;
            color: #999;
          }
        }

        .comment-content-right-header-right {
          .comment-content-right-header-right-like {
            font-size: 14px;
            cursor: pointer;

            .icon {
              //  这里可以添加喜欢的图标样式，例如：
              //  &::before {
              //    content: '\2665';
              //  }
            }
          }
        }
      }

      .comment-content-right-center {
        p {
          font-size: 14px;
          line-height: 1.5;
          // 设置显示行数
          overflow: hidden;
          text-overflow: ellipsis;
          display: -webkit-box;
          -webkit-box-orient: vertical;
          -webkit-line-clamp: 5;
          line-clamp: 5;
        }
      }
    }
  }

  .comment-footer {
    position: fixed;
    bottom: 0;
    width: 30%;
    display: flex;
    padding: 10px;
    box-sizing: border-box;

    .comment-footer-input {
      flex: 1;
      margin-right: 10px;

      input {
        width: 100%;
        padding: 8px;
        border: 1px solid #ccc;
        border-radius: 5px;
      }
    }

    .comment-footer-button {
      button {
        padding: 8px 15px;
        background-color: #4caf50;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
      }
    }
  }
}
</style>
