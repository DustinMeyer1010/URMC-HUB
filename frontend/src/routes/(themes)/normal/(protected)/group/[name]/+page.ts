export const prerender = true;

export const load = ({params}) => {

    return {
        name: params.name
    }
}



