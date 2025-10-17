# 🎵 Groupie Tracker

A dynamic web application that displays information about music bands and artists by consuming and manipulating a RESTful API. Features include advanced filtering, search functionality, and interactive geolocation mapping of concert venues.

![Go](https://img.shields.io/badge/Go-89.9%25-00ADD8?logo=go)
![CSS](https://img.shields.io/badge/CSS-5.9%25-1572B6?logo=css3)
![Docker](https://img.shields.io/badge/Docker-2.7%25-2496ED?logo=docker)

## 🎯 Features

### 📊 Data Visualization
- **Artist Information**: Display bands with images, members, inception year, and first album date
- **Concert Locations**: Show last and upcoming concert locations
- **Concert Dates**: Display past and upcoming concert dates
- **Relations**: Link artists with their concert locations and dates

### 🔍 Advanced Filtering
- Filter by **creation date range** (year)
- Filter by **first album date range**
- Filter by **number of members** (1-8)
- Filter by **concert location**
- Real-time filtering with form submissions

### 🔎 Search Functionality
- Search by artist/band name
- Search by member names
- Search by creation year
- Search by first album date
- Intelligent search with partial matching

### 🗺️ Geolocation
- MapQuest API integration
- Display concert locations on interactive maps
- Geographic coordinates visualization

## 🛠️ Technology Stack

### **Backend**
![Go](https://img.shields.io/badge/Go_1.21.3-00ADD8?style=for-the-badge&logo=go&logoColor=white)

- **Language**: Go (Golang) 1.21.3
- **Standard Library Only**: Pure Go implementation without external dependencies
- **Template Engine**: `text/template` for server-side rendering
- **HTTP Server**: Native Go `net/http`

### **Frontend**
- **HTML5**: Semantic markup and responsive design
- **CSS3**: Modern styling and animations
- **JavaScript**: Client-side interactions

### **APIs**
- **Groupie Trackers API**: Artist, location, date, and relation data
- **MapQuest API**: Geolocation and mapping services

### **DevOps**
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
- Docker containerization with Alpine Linux
- Automated deployment scripts

## 🚀 Quick Start

**Option 1: Run Locally**
```bash
# Clone the repository
git clone https://github.com/mamadbah2/groupie-tracker.git
cd groupie-tracker

# Run the application
go run ./cmd/web/.

# Access at http://0.0.0.0:1111
```

**Option 2: Docker Deployment**
```bash
# Build the Docker image
docker image build -f Dockerfile -t groupie-tracker .

# Run the container
docker run --detach -p 1111:1111 groupie-tracker

# Access at http://localhost:1111
```

## 📁 Project Structure

```
groupie-tracker/
├── cmd/web/              # Application entry point
│   └── run.go           # Server configuration
├── internal/
│   ├── handlers/        # HTTP handlers
│   │   ├── home.go     # Home page & filtering
│   │   ├── search.go   # Search functionality
│   │   └── repository.go # API data fetching
│   ├── models/          # Data structures
│   │   ├── artists.go  # Artist model
│   │   ├── template.go # Template data
│   │   └── map.go      # Geolocation models
│   └── pkg/             # Business logic
│       ├── filter.go   # Filtering algorithms
│       └── research.go # Search algorithms
├── templates/           # HTML templates
├── assets/             # Static files (CSS, images)
├── config/             # Application configuration
├── Dockerfile
└── go.mod
```

## 🎓 Skills Acquired

### **API Integration & Data Manipulation**
✅ **RESTful API Consumption** - Fetching and parsing JSON data  
✅ **Data Transformation** - Manipulating API responses  
✅ **Third-party API Integration** - MapQuest geolocation  
✅ **Error Handling** - Robust API error management

### **Backend Development**
✅ **Go Programming** - Standard library mastery  
✅ **HTTP Server** - Request/response handling  
✅ **Template Rendering** - Dynamic HTML generation  
✅ **Routing** - Custom HTTP routing logic  
✅ **Session Management** - State handling

### **Algorithm Design**
✅ **Filtering Algorithms** - Multi-criteria filtering  
✅ **Search Algorithms** - Partial and exact matching  
✅ **Data Processing** - Sorting and aggregation  
✅ **Performance Optimization** - Efficient data handling

### **Frontend Development**
✅ **Responsive Design** - Mobile-friendly interfaces  
✅ **Form Handling** - User input validation  
✅ **Dynamic UI** - Interactive data visualization  
✅ **Map Integration** - Geolocation display

### **Software Engineering**
✅ **Clean Architecture** - Separation of concerns  
✅ **Error Handling** - Comprehensive error management  
✅ **Code Organization** - Modular structure  
✅ **Testing** - Unit testing practices

### **DevOps**
✅ **Docker** - Containerization and deployment  
✅ **Git** - Version control and collaboration

## 🔧 Key Technical Achievements

1. **Pure Go Implementation** - Built entirely with standard library
2. **Advanced Filtering System** - Multi-parameter filtering with date ranges
3. **Intelligent Search** - Partial matching across multiple fields
4. **API Integration** - Consuming and manipulating external data sources
5. **Geolocation Mapping** - Geographic visualization of concert locations
6. **Robust Error Handling** - Graceful error management across all endpoints
7. **Template Caching** - Performance optimization for template rendering

## 📸 API Structure

The application consumes data from: `https://groupietrackers.herokuapp.com/api/`

**Endpoints:**
- `/artists` - Band and artist information
- `/locations` - Concert locations
- `/dates` - Concert dates
- `/relation` - Links between artists, locations, and dates

## 👥 Authors

- **[mamadbah2](https://github.com/mamadbah2)**

## 📝 License

Open source project - Available for educational purposes.

---

**Built with ❤️ at Zone01 Dakar**
