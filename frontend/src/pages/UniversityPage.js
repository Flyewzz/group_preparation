import React from "react";
import ListContainer from "../components/common/ListContainer"
import Subject from "../components/subjects/Subject";
import Filter from "../components/subjects/Filter";

const data = [
  {
    name: 'Физика',
    files: 2,
  },
  {
    name: 'Математический анализ',
    files: 2,
  },
  {
    name: 'Дискретная математика',
    files: 2,
  },
  {
    name: 'Экология',
    files: 2,
  },
  {
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
    name: 'Физика',
    files: 2,
  },
  {
    name: 'Математический анализ',
    files: 2,
  },
  {
    name: 'Дискретная математика',
    files: 2,
  },
  {
    name: 'Экология',
    files: 2,
  },
  {
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
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
