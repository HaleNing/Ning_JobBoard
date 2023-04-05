# Ning_JobBoard [Backend deploy on Render free plan, and still working on FrontEnd] 😎

> why this name?
> you may be curious,'Ning' in my personal name, in Germany, 'Xing' is a famous website for user find a new
> Job chances.
> more like Germany's linkedin.
> This project is want to create a job Board.
> users can Search for new Job on it.
>
> So, inspire by 'Xing'. I named this project as 'Ning'.
> THX for Render free plan ,you can deploy little project with free.

## Capturing Requirements 🤔

- As a visitor
    - I want to check all jobs in this job board website,and also can filter some conditional information(country,is
      Remote?,title) to check more speically job chances.
- As a logged in user
    - I still can do everything like visitor,but can post new job information to this Job Board website.or edit old job
      information created by this user.
- As administer
    - administer can do something like user,but as administer,it can edit everything in this website,like illegal
      message post by some bad user,administer can delete it.

## Tech stack 🧰

- Redis
- PGSQL
- Golang
- React
- Next.js
- Framework:
  - web: [fiber](https://github.com/gofiber/fiber)
  - orm: [ent](https://github.com/ent/ent)
