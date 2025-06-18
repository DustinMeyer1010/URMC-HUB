import { useState } from 'react';
import SearchBox from '../components/SearchBox';
import HomeStyles from './Home.module.css';
import LinkCard from '../components/LinkCard';
import Paging from '../components/Paging';

const Home = () => {

  const [searchValue, setSearchValue] = useState('');

    const handleEnter = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key != 'Enter') {
          return;
        }
        return
    }

    const handleClick = () => {
        return
    }

    const handleChange = (newValue: string) => {
        setSearchValue(newValue);
    }


    return (
    <div>
      <div className={HomeStyles.links_container}>
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
        <LinkCard />
      </div>
      <div className={HomeStyles.control_container}>
        <SearchBox
        value={searchValue}
        onChange={handleChange}
        onKeyDown={handleEnter}
        onClick={handleClick}
        />
        <Paging />
      </div>
    </div>
    );
  }


  export default Home;