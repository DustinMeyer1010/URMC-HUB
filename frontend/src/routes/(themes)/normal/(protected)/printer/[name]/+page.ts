export const prerender = true;

export const load = async ({params}) => {
    return {name: params.name}

}