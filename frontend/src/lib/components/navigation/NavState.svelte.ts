export class NavStateClass {
    ShowMenu: boolean  = $state(false)
    DropDownButton: HTMLButtonElement | undefined = $state(undefined)
    Width: number = $state(0)
    Compact: boolean = $derived(this.Width <= 500)

    HandleResize = () =>  {
            this.Width = window.innerWidth
        }


    ResizeSetup = () => {
        this.Width = window.innerWidth



        window.addEventListener("resize", this.HandleResize)

        document.addEventListener("click", (e) => {
            if (!this.DropDownButton?.contains(e.target as Node | null)) {
                this.ShowMenu = false
            }
        })
    }

    OpenMenu = () => {
        this.ShowMenu = !this.ShowMenu
    }

}
