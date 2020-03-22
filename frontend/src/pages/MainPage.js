import React from "react";
import University from "../components/universities/University";
import ListContainer from "../components/common/ListContainer"

const data = [
  {
    id: 0,
    short: 'МГТУ им. Баумана',
    full: 'Московский государственный технический университет имени Баумана'
  },
  {
    id: 1,
    short: 'ВШЭ',
    full: 'Высшая школа экономика'
  },
  {
    id: 2,
    short: 'МГУ им. Ломоносова',
    full: 'Московский государственный университет имени Ломоносова'
  },
  {
    id: 3,
    short: 'ГУУ',
    full: 'Государственный университет управления'
  },
  {
    id: 4,
    short: 'МГТУ им. Баумана',
    full: 'Московский государственный технический университет имени Баумана'
  },
  {
    id: 5,
    short: 'ВШЭ',
    full: 'Высшая школа экономика'
  },
  {
    id: 6,
    short: 'МГУ им. Ломоносова',
    full: 'Московский государственный университет имени Ломоносова'
  },
  {
    id: 7,
    short: 'ГУУ',
    full: 'Государственный университет управления'
  },
  {
    id: 8,
    short: 'МГУ им. Ломоносова',
    full: 'Московский государственный университет имени Ломоносова'
  },
];

function MainPage() {
  return (
    <ListContainer title={'Университеты'}
                   items={data.map((value) =>
                     <University key={value.id} university={value}/>
                   )}>
    </ListContainer>
  );
}

export default MainPage;
