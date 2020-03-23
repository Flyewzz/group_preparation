import React from "react";
import ListContainer from "../components/common/ListContainer"
import Subject from "../components/subjects/Subject";
import Filter from "../components/subjects/Filter";

const data = [
  {
    id: 1,
    name: 'Физика',
    files: 2,
  },
  {
    id: 2,
    name: 'Математический анализ',
    files: 2,
  },
  {
    id: 3,
    name: 'Дискретная математика',
    files: 2,
  },
  {
    id: 4,
    name: 'Экология',
    files: 2,
  },
  {
    id: 5,
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
    id: 6,
    name: 'Физика',
    files: 2,
  },
  {
    id: 7,
    name: 'Математический анализ',
    files: 2,
  },
  {
    id: 8,
    name: 'Дискретная математика',
    files: 2,
  },
  {
    id: 9,
    name: 'Экология',
    files: 2,
  },
  {
    id: 10,
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
    id: 11,
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
    id: 12,
    name: 'Аналитическая геометрия',
    files: 2,
  },
];

function UniversityPage() {
  return (
    <ListContainer title={'Предметы'}
                   subheader={<Filter/>}
                   items={data.map((value) =>
                     <Subject subject={value}/>
                   )}>
    </ListContainer>
  );
}

export default UniversityPage;
