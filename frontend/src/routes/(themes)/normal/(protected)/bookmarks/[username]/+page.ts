export const prerender = true;

export const load = ({params}) => {

    return {
        username: params.username
    }

}