import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/usage')({
  component: () => (
    <div>
      <h2>用量报表</h2>
      <p style={{ color: 'var(--td-text-color-secondary)' }}>
        GPU 计算时长、Token 消耗量、知识库调用次数 — 开发中
      </p>
    </div>
  ),
})
