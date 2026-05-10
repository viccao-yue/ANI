import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/settings/api-keys')({
  component: () => (
    <div>
      <h2>API Key 管理</h2>
      <p style={{ color: 'var(--td-text-color-secondary)' }}>创建和管理 API Key — 开发中</p>
    </div>
  ),
})
