import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/registry')({
  component: () => (
    <div>
      <h2>容器镜像仓库</h2>
      <p style={{ color: 'var(--td-text-color-secondary)' }}>
        Harbor API 封装 — 开发中
      </p>
    </div>
  ),
})
