import LinkCardStyle from "../components/LinkCard.module.css";
import image from "../assets/bac.jpg";

type LinkCard = {
    title?: string;
    link?: string;
    description?: string;
    imagePath?: string;
    favorite?: boolean;
}

console.log(image)


const LinkCard: React.FC<LinkCard> = ({title= "URMC-HUB", link = "", description = "This is going to be my test input", imagePath = image, favorite = true}) => {
    return (
        <a className={LinkCardStyle.a} href={link} target="_blank">
            <div className={LinkCardStyle.link_container}>

            <img className={LinkCardStyle.link_image} src={imagePath} alt={title} />
                <div className={LinkCardStyle.content_container}>
                    <h2>{title}</h2>
                    <p>{description}</p>
                </div>


                {favorite && <span>⭐</span>}
            </div>
        </a>
    )
}

export default LinkCard;