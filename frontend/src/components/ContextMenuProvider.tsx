import React, { createContext, useState, useContext, type ReactNode } from "react";
import ContextMenuStyles from "../styles/providers/ContextMenu.module.css"
type MenuState = {
  visible: boolean;
  x: number;
  y: number;
  content: ReactNode;
};

type ContextMenuContextType = {
  openMenu: (x: number, y: number, content: ReactNode) => void;
  closeMenu: () => void;
};

const ContextMenuContext = createContext<ContextMenuContextType | undefined>(undefined);

export const ContextMenuProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [menu, setMenu] = useState<MenuState>({ visible: false, x: 0, y: 0, content: null });

  const openMenu = (x: number, y: number, content: ReactNode) => {
    setMenu({ visible: true, x, y, content });
  };

  const closeMenu = () => {
    setMenu((m) => ({ ...m, visible: false }));
  };

  // Close menu on any click
  React.useEffect(() => {
    const close = () => closeMenu();
    document.addEventListener("click", close);
    return () => document.removeEventListener("click", close);
  }, []);

  return (
    <ContextMenuContext.Provider value={{ openMenu, closeMenu }}>
      {children}
      {menu.visible && (
        <div className={ContextMenuStyles.default} style={{top: menu.y, left: menu.x}}>
          {menu.content}
        </div>
      )}
    </ContextMenuContext.Provider>
  );
};

export const useContextMenu = () => {
  const ctx = useContext(ContextMenuContext);
  if (!ctx) throw new Error("useContextMenu must be used within a ContextMenuProvider");
  return ctx;
};