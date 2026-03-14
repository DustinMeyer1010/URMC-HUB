export const prerender = true;

export const load = ({params}) => {
    
    return {
        drive_name: decodeURIComponent(params.name)
    }
}