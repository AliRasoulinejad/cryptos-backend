package app

import (
	"gorm.io/gorm"
)

// AppName is the app's name .
const AppName string = "cryptos-backend"

const asciiArt = `
        CCCCCCCCCCCCC                                                                       tttt                                            
     CCC::::::::::::C                                                                    ttt:::t                                            
   CC:::::::::::::::C                                                                    t:::::t                                            
  C:::::CCCCCCCC::::C                                                                    t:::::t                                            
 C:::::C       CCCCCCrrrrr   rrrrrrrrryyyyyyy           yyyyyyyppppp   ppppppppp   ttttttt:::::ttttttt       ooooooooooo       ssssssssss   
C:::::C              r::::rrr:::::::::ry:::::y         y:::::y p::::ppp:::::::::p  t:::::::::::::::::t     oo:::::::::::oo   ss::::::::::s  
C:::::C              r:::::::::::::::::ry:::::y       y:::::y  p:::::::::::::::::p t:::::::::::::::::t    o:::::::::::::::oss:::::::::::::s 
C:::::C              rr::::::rrrrr::::::ry:::::y     y:::::y   pp::::::ppppp::::::ptttttt:::::::tttttt    o:::::ooooo:::::os::::::ssss:::::s
C:::::C               r:::::r     r:::::r y:::::y   y:::::y     p:::::p     p:::::p      t:::::t          o::::o     o::::o s:::::s  ssssss 
C:::::C               r:::::r     rrrrrrr  y:::::y y:::::y      p:::::p     p:::::p      t:::::t          o::::o     o::::o   s::::::s      
C:::::C               r:::::r               y:::::y:::::y       p:::::p     p:::::p      t:::::t          o::::o     o::::o      s::::::s   
 C:::::C       CCCCCC r:::::r                y:::::::::y        p:::::p    p::::::p      t:::::t    tttttto::::o     o::::ossssss   s:::::s 
  C:::::CCCCCCCC::::C r:::::r                 y:::::::y         p:::::ppppp:::::::p      t::::::tttt:::::to:::::ooooo:::::os:::::ssss::::::s
   CC:::::::::::::::C r:::::r                  y:::::y          p::::::::::::::::p       tt::::::::::::::to:::::::::::::::os::::::::::::::s 
     CCC::::::::::::C r:::::r                 y:::::y           p::::::::::::::pp          tt:::::::::::tt oo:::::::::::oo  s:::::::::::ss  
        CCCCCCCCCCCCC rrrrrrr                y:::::y            p::::::pppppppp              ttttttttttt     ooooooooooo     sssssssssss    
                                            y:::::y             p:::::p                                                                     
                                           y:::::y              p:::::p                                                                     
                                          y:::::y              p:::::::p                                                                    
                                         y:::::y               p:::::::p                                                                    
                                        yyyyyyy                p:::::::p                                                                    
                                                               ppppppppp
`

type Application struct {
	DB           *gorm.DB
	Repositories *Repositories
}

func Banner() string {
	return asciiArt
}
