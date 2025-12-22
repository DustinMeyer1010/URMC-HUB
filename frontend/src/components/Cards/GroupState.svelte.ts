import { Group } from "@t/group";

export class GroupStateClass {
    CurrentGroup: Group.CardInfo = $state(Group.EMPTY_GROUP)
    Idx: number = $state(0)
    CopyText: string = $derived(`Name: ${this.CurrentGroup.name}\nInformation: ${this.CurrentGroup.information !== "" ? this.CurrentGroup.information : "NA"}\nDescription: ${this.CurrentGroup.description !== "" ? this.CurrentGroup.description : "NA"}\nOU: ${this.CurrentGroup.ou}`)
    constructor(group: Group.CardInfo, idx: number) {
        this.CurrentGroup = group
        this.Idx = idx
    }
}