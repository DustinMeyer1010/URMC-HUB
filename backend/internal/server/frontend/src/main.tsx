import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { AuthProvider } from './components/Authentication.tsx'
import { ContextMenuProvider } from './components/ContextMenuProvider.tsx'

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <AuthProvider>
            <ContextMenuProvider>
                <App />
            </ContextMenuProvider>
        </AuthProvider>
    </StrictMode>,
)
